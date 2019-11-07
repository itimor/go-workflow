package sys

import (
	"iris-ticket/backend/controllers/common"
	models "iris-ticket/backend/models/common"
	"iris-ticket/backend/models/sys"
	"iris-ticket/backend/pkg/convert"

	"github.com/gin-gonic/gin"
	"github.com/kataras/iris"
)

type Menu struct{}

// 分页数据
func (Menu) List(ctx iris.Context) {
	page := common.GetPageIndex(ctx)
	limit := common.GetPageLimit(ctx)
	sort := common.GetPageSort(ctx)
	key := common.GetPageKey(ctx)
	menuType := common.GetQueryToUint(ctx, "type")
	parent_id := common.GetQueryToUint64(ctx, "parent_id")
	var whereOrder []models.PageWhereOrder
	order := "ID DESC"
	if len(sort) >= 2 {
		orderType := sort[0:1]
		order = sort[1:len(sort)]
		if orderType == "+" {
			order += " ASC"
		} else {
			order += " DESC"
		}
	}
	whereOrder = append(whereOrder, models.PageWhereOrder{Order: order})
	if key != "" {
		v := "%" + key + "%"
		var arr []interface{}
		arr = append(arr, v)
		arr = append(arr, v)
		whereOrder = append(whereOrder, models.PageWhereOrder{Where: "name like ? or code like ?", Value: arr})
	}
	if menuType > 0 {
		var arr []interface{}
		arr = append(arr, menuType)
		whereOrder = append(whereOrder, models.PageWhereOrder{Where: "menu_type = ?", Value: arr})
	}
	if parent_id > 0 {
		var arr []interface{}
		arr = append(arr, parent_id)
		whereOrder = append(whereOrder, models.PageWhereOrder{Where: "parent_id = ?", Value: arr})
	}
	var total uint64
	list := []sys.Menu{}
	err := models.GetPage(&sys.Menu{}, &sys.Menu{}, &list, page, limit, &total, whereOrder...)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccessPage(ctx, total, &list)
}

// 详情
func (Menu) Detail(ctx iris.Context) {
	id := common.GetQueryToUint64(ctx, "id")
	var menu sys.Menu
	where := sys.Menu{}
	where.ID = id
	_, err := models.First(&where, &menu)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccess(ctx, &menu)
}

// 更新
func (Menu) Update(ctx iris.Context) {
	model := sys.Menu{}
	err := ctx.ReadJSON(&model)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	err = models.Save(&model)
	if err != nil {
		common.ResFail(ctx, "操作失败")
		return
	}
	common.ResSuccessMsg(ctx)
}

//新增
func (Menu) Create(ctx iris.Context) {
	menu := sys.Menu{}
	err := ctx.ReadJSON(&menu)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	err = models.Create(&menu)
	if err != nil {
		common.ResFail(ctx, "操作失败")
		return
	}
	go InitMenu(menu)
	common.ResSuccess(ctx, gin.H{"id": menu.ID})
}

// 删除数据
func (Menu) Delete(ctx iris.Context) {
	var ids []uint64
	err := ctx.ReadJSON(&ids)
	if err != nil || len(ids) == 0 {
		common.ResErrSrv(ctx, err)
		return
	}
	menu := sys.Menu{}
	err = menu.Delete(ids)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccessMsg(ctx)
}

// 所有菜单
func (Menu) AllMenu(ctx iris.Context) {
	var menus []sys.Menu
	err := models.Find(&sys.Menu{}, &menus, "parent_id asc", "sequence asc")
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccess(ctx, &menus)
}

// 新增菜单后自动添加菜单下的常规操作
func InitMenu(model sys.Menu) {
	if model.MenuType != 2 {
		return
	}
	add := sys.Menu{Status: 1, ParentID: model.ID, URL: model.URL + "/create", Name: "新增", Sequence: 1, MenuType: 3, Code: model.Code + "Add", OperateType: "add"}
	models.Create(&add)
	del := sys.Menu{Status: 1, ParentID: model.ID, URL: model.URL + "/delete", Name: "删除", Sequence: 2, MenuType: 3, Code: model.Code + "Del", OperateType: "del"}
	models.Create(&del)
	view := sys.Menu{Status: 1, ParentID: model.ID, URL: model.URL + "/detail", Name: "查看", Sequence: 3, MenuType: 3, Code: model.Code + "View", OperateType: "view"}
	models.Create(&view)
	update := sys.Menu{Status: 1, ParentID: model.ID, URL: model.URL + "/update", Name: "编辑", Sequence: 4, MenuType: 3, Code: model.Code + "Update", OperateType: "update"}
	models.Create(&update)
	list := sys.Menu{Status: 1, ParentID: model.ID, URL: model.URL + "/list", Name: "分页api", Sequence: 5, MenuType: 3, Code: model.Code + "List", OperateType: "list"}
	models.Create(&list)
}

// 获取菜单有权限的操作列表
func (Menu) MenuButtonList(ctx iris.Context) {
	// 用户ID
	uid, err := ctx.Values().GetUint64("auth_user_id")
	if err != nil {
		common.ResFailCode(ctx, "token 无效", 50008)
		return
	}
	userID := convert.ToUint64(uid)
	menuCode := common.GetQueryToStr(ctx, "menucode")
	if userID == 0 || menuCode == "" {
		common.ResFail(ctx, "err")
		return
	}
	btnList := []string{}
	if userID == common.SUPER_ADMIN_ID {
		//管理员
		btnList = append(btnList, "add")
		btnList = append(btnList, "del")
		btnList = append(btnList, "view")
		btnList = append(btnList, "update")
		btnList = append(btnList, "setrolemenu")
		btnList = append(btnList, "setadminrole")
	} else {
		menu := sys.Menu{}
		err := menu.GetMenuButton(userID, menuCode, &btnList)
		if err != nil {
			common.ResErrSrv(ctx, err)
			return
		}
	}
	common.ResSuccess(ctx, &btnList)
}
