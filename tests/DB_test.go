package tests

import (
	"haocean/health-enforcement/app/admin/model/system"
	"haocean/health-enforcement/config"
	"haocean/health-enforcement/pkg/db"
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
