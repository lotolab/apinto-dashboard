package random_controller

import (
	"github.com/eolinker/apinto-dashboard/common"
	"github.com/eolinker/apinto-dashboard/controller"
	"github.com/eolinker/apinto-dashboard/modules/random"
	"github.com/eolinker/eosc/common/bean"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type randomController struct {
	randomService random.IRandomService
}

func RegisterRandomRouter(router gin.IRoutes) {
	r := &randomController{}
	bean.Autowired(&r.randomService)

	router.GET("/random/:template/id", r.get)
}

func (r *randomController) get(ginCtx *gin.Context) {
	template := ginCtx.Param("template")
	randomStr := r.randomService.RandomStr(template)
	m := common.Map[string, interface{}]{}

	m["id"] = strings.ToLower(randomStr)

	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(m))
}