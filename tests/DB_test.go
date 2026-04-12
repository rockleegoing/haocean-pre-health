package tests

import (
	"ruoyi-go/app/admin/model/system"
	"ruoyi-go/config"
	"ruoyi-go/pkg/db"
	"testing"
)

func TestDB(t *testing.T) {
	config.InitAppConfig(*configFile)
	var list []system.SysUser
	var total int64
	var db = db.Exe().Model(system.SysUser{})

	if err := db.Count(&total).Error; err != nil {
		println(err)
	}
	db.Find(&list)
	println(len(list))
	println(total)
}
