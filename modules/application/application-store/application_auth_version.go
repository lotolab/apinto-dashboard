package application_store

import (
	"encoding/json"
	"github.com/eolinker/apinto-dashboard/modules/application/application-entry"
	"github.com/eolinker/apinto-dashboard/modules/base/version-entry"
	"github.com/eolinker/apinto-dashboard/store"
)

type IApplicationAuthVersionStore interface {
	store.IBaseStore[application_entry.ApplicationAuthVersion]
}

type applicationAuthVersionKindHandler struct {
}

func (s *applicationAuthVersionKindHandler) Kind() string {
	return "application_auth"
}

func (s *applicationAuthVersionKindHandler) Encode(sv *application_entry.ApplicationAuthVersion) *version_entry.Version {

	v := new(version_entry.Version)
	v.Id = sv.Id
	v.Kind = s.Kind()
	v.Target = sv.ApplicationAuthID
	v.Operator = sv.Operator
	v.NamespaceID = sv.NamespaceID
	v.CreateTime = sv.CreateTime
	bytes, _ := json.Marshal(sv.ApplicationAuthVersionConfig)
	v.Data = bytes

	return v
}

func (s *applicationAuthVersionKindHandler) Decode(v *version_entry.Version) *application_entry.ApplicationAuthVersion {
	sv := new(application_entry.ApplicationAuthVersion)
	sv.Id = v.Id
	sv.ApplicationAuthID = v.Target
	sv.Operator = v.Operator
	sv.NamespaceID = v.NamespaceID
	sv.CreateTime = v.CreateTime
	_ = json.Unmarshal(v.Data, &sv.ApplicationAuthVersionConfig)

	return sv
}

func newApplicationAuthVersionStore(db store.IDB) IApplicationAuthVersionStore {
	var h store.BaseKindHandler[application_entry.ApplicationAuthVersion, version_entry.Version] = new(applicationAuthVersionKindHandler)
	return store.CreateBaseKindStore(h, db)
}