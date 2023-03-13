package store

import (
	"context"
	"github.com/eolinker/apinto-dashboard/entry"
	"gorm.io/gorm"
)

type IGlobalVariableStore interface {
	IBaseStore[entry.Variables]
	GetList(ctx context.Context, pageNum, pageSize, namespaceID int, searchName string) ([]*entry.Variables, int, error)
	GetGlobalVariableIDByKey(ctx context.Context, namespaceID int, key string) (*entry.Variables, error)
	GetGlobalVariableByKeys(ctx context.Context, namespaceID int, keys []string) ([]*entry.Variables, error)
}

type globalVariableStore struct {
	*baseStore[entry.Variables]
}

func newGlobalVariableStore(db IDB) IGlobalVariableStore {
	return &globalVariableStore{baseStore: createStore[entry.Variables](db)}
}

func (g *globalVariableStore) GetList(ctx context.Context, pageNum, pageSize, namespaceID int, searchName string) ([]*entry.Variables, int, error) {
	variables := make([]*entry.Variables, 0)
	db := g.DB(ctx).Where("namespace = ?", namespaceID)

	count := int64(0)
	if searchName != "" {
		db = db.Where("`key` like ?", "%"+searchName+"%")
	}

	if pageNum > 0 && pageSize > 0 {
		err := db.Model(variables).Count(&count).Order("create_time asc").Limit(pageSize).Offset(entry.PageIndex(pageNum, pageSize)).Find(&variables).Error
		if err != nil {
			return nil, 0, err
		}
	} else {
		err := db.Order("create_time asc").Find(&variables).Error
		if err != nil {
			return nil, 0, err
		}
	}
	return variables, int(count), nil
}

func (g *globalVariableStore) GetGlobalVariableIDByKey(ctx context.Context, namespaceID int, key string) (*entry.Variables, error) {
	variable := &entry.Variables{}
	if err := g.DB(ctx).Where("namespace = ? AND `key` = ?", namespaceID, key).Take(variable).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return variable, nil
}

func (g *globalVariableStore) GetGlobalVariableByKeys(ctx context.Context, namespaceID int, keys []string) ([]*entry.Variables, error) {
	variables := make([]*entry.Variables, 0, len(keys))
	if err := g.DB(ctx).Where("namespace = ? AND `key` in ?", namespaceID, keys).Find(&variables).Error; err != nil {
		return nil, err
	}
	return variables, nil
}
