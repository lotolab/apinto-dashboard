package driver

import (
	"fmt"
	v1 "github.com/eolinker/apinto-dashboard/client/v1"
	"github.com/eolinker/apinto-dashboard/common"
	"github.com/eolinker/apinto-dashboard/enum"
	"github.com/eolinker/apinto-dashboard/modules/api"
	"github.com/eolinker/apinto-dashboard/modules/api/api-dto"
	api_entry "github.com/eolinker/apinto-dashboard/modules/api/api-entry"
	"net/textproto"
	"strings"
)

type apiHTTP struct {
	apintoDriverName string
}

func (a *apiHTTP) CheckInput(input *api_dto.APIInfo) error {
	for _, m := range input.Method {
		switch m {
		case enum.MethodGET, enum.MethodPOST, enum.MethodPUT, enum.MethodDELETE, enum.MethodPATCH, enum.MethodHEAD, enum.MethodOPTIONS:
		default:
			return fmt.Errorf("method %s is illegal. ", m)
		}
	}

	return checkInput(input)
}

func (a *apiHTTP) ToApinto(name, desc string, disable bool, method []string, requestPath, requestPathLabel, proxyPath, serviceName string, timeout, retry int, enableWebsocket bool, match []*api_entry.MatchConf, header []*api_entry.ProxyHeader, templateUUID string) *v1.RouterConfig {

	rewriteHeaders := make(map[string]string)
	for _, ph := range header {
		//格式化header Key
		ph.Key = textproto.CanonicalMIMEHeaderKey(ph.Key)
		switch ph.OptType {
		case enum.HeaderOptTypeAdd:
			rewriteHeaders[ph.Key] = ph.Value
		case enum.HeaderOptTypeDelete:
			rewriteHeaders[ph.Key] = ""
		}
	}

	rewritePlugin := v1.PluginProxyRewriteV2Config{
		NotMatchErr: true,
		HostRewrite: false,
		Headers:     rewriteHeaders,
	}
	//若请求路径包含restful参数
	if common.IsRestfulPath(requestPath) {
		rewritePlugin.PathType = "regex" //正则替换

		//如果转发路径包含restful参数
		if common.IsRestfulPath(proxyPath) {
			proxyPath = formatProxyPath(requestPathLabel, proxyPath)
		}
		rewritePlugin.RegexPath = []*v1.RegexPath{
			{
				RegexPathMatch:   fmt.Sprintf("^%s$", common.ReplaceRestfulPath(requestPath, apintoRestfulRegexp)),
				RegexPathReplace: proxyPath,
			},
		}
		requestPath = fmt.Sprintf("~=^%s$", common.ReplaceRestfulPath(requestPath, apintoRestfulRegexp))
	} else {
		rewritePlugin.PathType = "prefix" //前缀替换
		rewritePlugin.PrefixPath = []*v1.PrefixPath{
			{
				PrefixPathMatch:   strings.TrimSuffix(requestPath, "*"),
				PrefixPathReplace: proxyPath,
			},
		}
	}

	rules := make([]*v1.RouterRule, 0, len(match))
	for _, m := range match {
		rule := &v1.RouterRule{
			Type:  m.Position,
			Name:  m.Key,
			Value: "",
		}

		if m.Position == enum.MatchPositionHeader {
			rule.Name = textproto.CanonicalMIMEHeaderKey(rule.Name)
		}

		switch m.MatchType {
		case enum.MatchTypeEqual:
			rule.Value = m.Pattern
		case enum.MatchTypePrefix:
			rule.Value = fmt.Sprintf("%s*", m.Pattern)
		case enum.MatchTypeSuffix:
			rule.Value = fmt.Sprintf("*%s", m.Pattern)
		case enum.MatchTypeSubstr:
			rule.Value = fmt.Sprintf("*%s*", m.Pattern)
		case enum.MatchTypeUneuqal:
			rule.Value = fmt.Sprintf("!=%s", m.Pattern)
		case enum.MatchTypeNull:
			rule.Value = "$"
		case enum.MatchTypeExist:
			rule.Value = "**"
		case enum.MatchTypeUnexist:
			rule.Value = "!"
		case enum.MatchTypeRegexp:
			rule.Value = fmt.Sprintf("~=%s", m.Pattern)
		case enum.MatchTypeRegexpG:
			rule.Value = fmt.Sprintf("~*=%s", m.Pattern)
		case enum.MatchTypeAny:
			rule.Value = "*"
		}

		rules = append(rules, rule)
	}

	templateID := ""
	if templateUUID != "" {
		templateID = fmt.Sprintf("%s@template", templateUUID)
	}
	return &v1.RouterConfig{
		Name:            name,
		Description:     desc,
		Driver:          a.apintoDriverName,
		Disable:         disable,
		Listen:          0,
		Method:          method,
		Host:            []string{},
		RequestPath:     requestPath,
		Rules:           rules,
		Service:         fmt.Sprintf("%s@service", serviceName),
		Template:        templateID,
		Retry:           retry,
		Timeout:         timeout,
		EnableWebsocket: enableWebsocket,
		Plugins: map[string]*v1.Plugin{
			"proxy_rewrite": { //插件名写死
				Disable: false,
				Config:  rewritePlugin,
			},
		},
	}
}

func CreateAPIHttp(apintoDriverName string) api.IAPIDriver {
	return &apiHTTP{apintoDriverName: apintoDriverName}
}
