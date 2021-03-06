package source

import (
	"RemindMe/model"
	"time"

	"RemindMe/global"
	"RemindMe/model/system"
	"github.com/gookit/color"

	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

var menus = []system.SysBaseMenu{
	{Model: models.Model{ID: 1, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1, Meta: system.Meta{Title: "仪表盘", Icon: "setting"}},
	{Model: models.Model{ID: 2, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7, Meta: system.Meta{Title: "关于我们", Icon: "info"}},
	{Model: models.Model{ID: 3, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: system.Meta{Title: "超级管理员", Icon: "user-solid"}},
	{Model: models.Model{ID: 4, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: system.Meta{Title: "角色管理", Icon: "s-custom"}},
	{Model: models.Model{ID: 5, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: system.Meta{Title: "菜单管理", Icon: "s-order", KeepAlive: true}},
	{Model: models.Model{ID: 6, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: system.Meta{Title: "API管理", Icon: "s-platform", KeepAlive: true}},
	{Model: models.Model{ID: 7, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: system.Meta{Title: "用户管理", Icon: "coordinate"}},
	{Model: models.Model{ID: 8, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: system.Meta{Title: "个人信息", Icon: "message-solid"}},
	{Model: models.Model{ID: 9, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 6, Meta: system.Meta{Title: "示例文件", Icon: "s-management"}},
	{Model: models.Model{ID: 10, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "excel", Name: "excel", Component: "view/example/excel/excel.vue", Sort: 4, Meta: system.Meta{Title: "excel导入导出", Icon: "s-marketing"}},
	{Model: models.Model{ID: 11, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: system.Meta{Title: "媒体库（上传下载）", Icon: "upload"}},
	{Model: models.Model{ID: 12, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: system.Meta{Title: "断点续传", Icon: "upload"}},
	{Model: models.Model{ID: 13, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: system.Meta{Title: "客户列表（资源示例）", Icon: "s-custom"}},
	{Model: models.Model{ID: 14, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: system.Meta{Title: "系统工具", Icon: "s-cooperation"}},
	{Model: models.Model{ID: 15, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: system.Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
	{Model: models.Model{ID: 16, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: system.Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true}},
	{Model: models.Model{ID: 17, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: system.Meta{Title: "系统配置", Icon: "s-operation"}},
	{Model: models.Model{ID: 18, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: system.Meta{Title: "字典管理", Icon: "notebook-2"}},
	{Model: models.Model{ID: 19, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: system.Meta{Title: "字典详情", Icon: "s-order"}},
	{Model: models.Model{ID: 20, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: system.Meta{Title: "操作历史", Icon: "time"}},
	{Model: models.Model{ID: 22, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, ParentId: "0", Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Hidden: false, Component: "/", Sort: 0, Meta: system.Meta{Title: "官方网站", Icon: "s-home"}},
	{Model: models.Model{ID: 23, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "view/system/state.vue", Sort: 6, Meta: system.Meta{Title: "服务器状态", Icon: "cloudy"}},
	{Model: models.Model{ID: 24, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, ParentId: "14", Path: "autoCodeAdmin", Name: "autoCodeAdmin", Hidden: false, Component: "view/systemTools/autoCodeAdmin/index.vue", Sort: 1, Meta: system.Meta{Title: "自动化代码管理", Icon: "s-finance"}},
	{Model: models.Model{ID: 25, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, MenuLevel: 0, ParentId: "14", Path: "autoCodeEdit/:id", Name: "autoCodeEdit", Hidden: true, Component: "view/systemTools/autoCode/index.vue", Sort: 0, Meta: system.Meta{Title: "自动化代码（复用）", Icon: "s-finance"}},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_base_menus 表数据初始化
func (m *menu) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 29}).Find(&[]system.SysBaseMenu{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_base_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_base_menus 表初始数据成功!")
		return nil
	})
}
