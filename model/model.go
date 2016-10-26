package model

import (
	"audit-rule/config"
	"github.com/caojia/go-orm"
)

type Model struct {
	m     *orm.ORM
	slave *orm.ORM
}

func NewModel() *Model {
	ret := &Model{
		m:     orm.NewORM(config.Instance().Db.Master),
		slave: orm.NewORM(config.Instance().Db.Slave),
	}

	return ret
}
