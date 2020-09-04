package service

import (
	"github.com/itchin/gin-demo/helper/orm"
	"github.com/itchin/gin-demo/model"
)

type tool struct{}

var Tool tool

func (*tool) List() (list []*model.Tool, err error) {
	md := new(model.Tool)
	engine := orm.Engine()
	err = engine.Table(md.TableName()).Select("*").Limit(10).Find(&list)
	if err != nil {
		return
	}
	if list == nil {
		list = []*model.Tool{}
	}
	return
}
