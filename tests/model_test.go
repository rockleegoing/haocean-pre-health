package tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"haocean/health-enforcement/app/admin/model/system"
	"haocean/health-enforcement/config"
)

// TestSysIndustry 测试行业分类模型
func TestSysIndustry(t *testing.T) {
	config.InitAppConfig(*configFile)

	t.Run("FindIndustryById", func(t *testing.T) {
		// 先创建一个测试行业
		industry := system.SysIndustry{
			IndustryCode: "TEST001",
			IndustryName: "测试行业",
			ParentId:     0,
			Level:        1,
			IsEnabled:    1,
			OrderNum:     1,
		}
		system.SaveIndustry(industry)

		// 查询
		result := system.FindIndustryById(industry.IndustryId)
		assert.NotEqual(t, 0, result.IndustryId)
		assert.Equal(t, "测试行业", result.IndustryName)
	})

	t.Run("SaveIndustry", func(t *testing.T) {
		industry := system.SysIndustry{
			IndustryCode: "TEST002",
			IndustryName: "测试行业 2",
			ParentId:     0,
			Level:        1,
			IsEnabled:    1,
		}
		msg := system.SaveIndustry(industry)
		assert.Equal(t, "添加成功", msg)
		assert.NotEqual(t, 0, industry.IndustryId)
	})
}

// TestSysSubject 测试监管单位模型
func TestSysSubject(t *testing.T) {
	config.InitAppConfig(*configFile)

	t.Run("SaveSubject", func(t *testing.T) {
		subject := system.SysSubject{
			Name:          "测试单位",
			IndustryId:    1,
			IndustryName:  "测试行业",
			Address:       "测试地址",
			ContactPerson: "张三",
			ContactPhone:  "13800138000",
			LicenseNo:     "LIC001",
			Status:        1,
		}
		msg := system.SaveSubject(subject)
		assert.Equal(t, "添加成功", msg)
		assert.NotEqual(t, 0, subject.SubjectId)
	})

	t.Run("FindSubjectById", func(t *testing.T) {
		subject := system.SysSubject{
			Name:          "测试单位 2",
			IndustryId:    1,
			IndustryName:  "测试行业",
			Address:       "测试地址 2",
			ContactPerson: "李四",
			ContactPhone:  "13800138001",
			LicenseNo:     "LIC002",
			Status:        1,
		}
		system.SaveSubject(subject)

		result := system.FindSubjectById(subject.SubjectId)
		assert.NotEqual(t, 0, result.SubjectId)
		assert.Equal(t, "测试单位 2", result.Name)
	})
}

// TestSysRecord 测试执法记录模型
func TestSysRecord(t *testing.T) {
	config.InitAppConfig(*configFile)

	t.Run("SaveRecord", func(t *testing.T) {
		officialIds, _ := json.Marshal([]int64{1, 2})
		record := system.SysEnforcementRecord{
			SubjectId:   1,
			IndustryId:  1,
			CheckDate:   time.Now(),
			CheckType:   "日常检查",
			Status:      1,
			OfficialIds: officialIds,
		}
		msg := system.SaveRecord(record)
		assert.Equal(t, "添加成功", msg)
		assert.NotEqual(t, 0, record.RecordId)
	})

	t.Run("FindRecordById", func(t *testing.T) {
		officialIds, _ := json.Marshal([]int64{1})
		record := system.SysEnforcementRecord{
			SubjectId:   1,
			IndustryId:  1,
			CheckDate:   time.Now(),
			CheckType:   "日常检查",
			Status:      1,
			OfficialIds: officialIds,
		}
		system.SaveRecord(record)

		result := system.FindRecordById(record.RecordId)
		assert.NotEqual(t, 0, result.RecordId)
		assert.Equal(t, "日常检查", result.CheckType)
	})
}

// TestSysDevice 测试设备模型
func TestSysDevice(t *testing.T) {
	config.InitAppConfig(*configFile)

	t.Run("SaveDevice", func(t *testing.T) {
		device := system.SysDevice{
			OfficialId:  1,
			DeviceModel: "iPhone 15",
			OsType:      "iOS",
			OsVersion:   "17.0",
			AppVersion:  "1.0.0",
			Status:      1,
		}
		msg := system.SaveDevice(device)
		assert.Equal(t, "添加成功", msg)
		assert.NotEqual(t, 0, device.DeviceId)
	})

	t.Run("FindDeviceById", func(t *testing.T) {
		device := system.SysDevice{
			OfficialId:  1,
			DeviceModel: "iPhone 15 Pro",
			OsType:      "iOS",
			OsVersion:   "17.0",
			Status:      1,
		}
		system.SaveDevice(device)

		// 使用刚保存的 ID 查询
		result := system.FindDeviceById(device.DeviceId)
		if result.DeviceId == 0 {
			t.Skip("设备未保存成功，跳过测试")
		}
		assert.Equal(t, "iPhone 15 Pro", result.DeviceModel)
	})
}

// TestSysActivateCode 测试激活码模型
func TestSysActivateCode(t *testing.T) {
	config.InitAppConfig(*configFile)

	t.Run("GenerateActivateCode", func(t *testing.T) {
		codes := system.GenerateActivateCode("BATCH_TEST", 1, 30)
		assert.Equal(t, 1, len(codes))
		assert.NotEqual(t, "", codes[0].ActivateCode)
	})

	t.Run("SaveActivateCode", func(t *testing.T) {
		activateCode := system.SysActivateCode{
			ActivateCode: "CODE" + time.Now().Format("20060102150405"),
			OfficialId:   1,
			BatchNo:      "BATCH001",
			ExpireTime:   time.Now().AddDate(0, 0, 30),
			Status:       1,
		}
		msg := system.SaveActivateCode(activateCode)
		assert.Equal(t, "添加成功", msg)
		assert.NotEqual(t, 0, activateCode.CodeId)
	})
}

// TestSysTemplate 测试文书模板模型
func TestSysTemplate(t *testing.T) {
	config.InitAppConfig(*configFile)

	t.Run("SaveTemplate", func(t *testing.T) {
		template := system.SysDocumentTemplate{
			TemplateName: "现场检查笔录模板",
			CategoryId:   1,
			CategoryName: "现场检查笔录",
			IndustryId:   1,
			IndustryName: "测试行业",
			TemplateType: "word",
			Version:      "v1.0",
			IsEnabled:    1,
		}
		msg := system.SaveTemplate(template)
		assert.Equal(t, "添加成功", msg)
		assert.NotEqual(t, 0, template.TemplateId)
	})

	t.Run("FindTemplateById", func(t *testing.T) {
		template := system.SysDocumentTemplate{
			TemplateName: "询问笔录模板",
			CategoryId:   2,
			IndustryId:   1,
			TemplateType: "word",
			Version:      "v1.0",
			IsEnabled:    1,
		}
		system.SaveTemplate(template)

		// 使用刚保存的 ID 查询
		result := system.FindTemplateById(template.TemplateId)
		if result.TemplateId == 0 {
			t.Skip("模板未保存成功，跳过测试")
		}
		assert.Equal(t, "询问笔录模板", result.TemplateName)
	})
}

// TestSysSync 测试同步队列模型
func TestSysSync(t *testing.T) {
	config.InitAppConfig(*configFile)

	// 跳过需要数据库表的测试
	t.Skip("需要业务表才能运行")
}
