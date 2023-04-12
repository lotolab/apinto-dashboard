package entry

import (
	"time"
)

type ModulePlugin struct {
	Id           int       `gorm:"column:id" json:"id"`
	UUID         string    `gorm:"column:uuid" json:"uuid"`
	Name         string    `gorm:"column:name" json:"name"`
	Version      string    `gorm:"column:version" json:"version"`
	Group        int       `gorm:"column:group" json:"group"`
	NavigationID string    `gorm:"column:navigation_id" json:"navigation_id"`
	CName        string    `gorm:"column:cname" json:"cname"`
	Resume       string    `gorm:"column:resume" json:"resume"`
	ICon         string    `gorm:"column:icon" json:"icon"`
	Type         int       `gorm:"column:type" json:"type"`
	Front        string    `gorm:"column:front" json:"front"`
	Driver       string    `gorm:"column:driver" json:"driver"`
	Details      []byte    `gorm:"column:details" json:"details"`
	Operator     int       `gorm:"column:operator" json:"operator"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
}

func (*ModulePlugin) TableName() string {
	return "module_plugin"
}

func (c *ModulePlugin) IdValue() int {
	return c.Id
}

type EnablePlugin struct {
	UUID   string
	Name   string
	Driver string
	Config []byte
	Define []byte
}

// EnabledModule 启用的导航模块信息
type EnabledModule struct {
	Name         string `gorm:"column:name" json:"name"`
	Title        string `gorm:"column:cname" json:"cname"`
	Type         int    `gorm:"column:type" json:"type"`
	NavigationID int    `gorm:"column:navigation" json:"navigation"`
	Front        string `gorm:"column:front" json:"front"`
}
