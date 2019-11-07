package sys

import (
	"iris-ticket/backend/controllers/common"
	"iris-ticket/backend/middleware/casbins"
	models "iris-ticket/backend/models/common"
	"iris-ticket/backend/models/sys"

	"github.com/gin-gonic/gin"
	"github.com/kataras/iris"
)

type Role struct{}

// 分页数据
func (Role) List(ctx iris.Context) {
	page := common.GetPageIndex(ctx)
	limit := common.GetPageLimit(ctx)
	sort := common.GetPageSort(ctx)
	key := common.GetPageKey(ctx)
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
		whereOrder = append(whereOrder, models.PageWhereOrder{Where: "name like ?", Value: arr})
	}
	if parent_id > 0 {
		var arr []interface{}
		arr = append(arr, parent_id)
		whereOrder = append(whereOrder, models.PageWhereOrder{Where: "parent_id = ?", Value: arr})
	}
	var total uint64
	list := []sys.Role{}
	err := models.GetPage(&sys.Role{}, &sys.Role{}, &list, page, limit, &total, whereOrder...)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccessPage(ctx, total, &list)
}

// 详情
func (Role) Detail(ctx iris.Context) {
	id := common.GetQueryToUint64(ctx, "id")
	var model sys.Role
	where := sys.Role{}
	where.ID = id
	_, err := models.First(&where, &model)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccess(ctx, &model)
}

// 更新
func (Role) Update(ctx iris.Context) {
	model := sys.Role{}
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
func (Role) Create(ctx iris.Context) {
	model := sys.Role{}
	err := ctx.ReadJSON(&model)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	err = models.Create(&model)
	if err != nil {
		common.ResFail(ctx, "操作失败")
		return
	}
	common.ResSuccess(ctx, gin.H{"id": model.ID})
}

// 删除数据
func (Role) Delete(ctx iris.Context) {
	var ids []uint64
	err := ctx.ReadJSON(&ids)
	if err != nil || len(ids) == 0 {
		common.ResErrSrv(ctx, err)
		return
	}
	role := sys.Role{}
	err = role.Delete(ids)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	go casbins.CsbinDeleteRole(ids)
	common.ResSuccessMsg(ctx)
}

// 所有角色
func (Role) AllRole(ctx iris.Context) {
	var list []sys.Role
	err := models.Find(&sys.Role{}, &list, "parent_id asc", "sequence asc")
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccess(ctx, &list)
}

// 获取角色下的菜单ID列表
func (Role) RoleMenuIDList(ctx iris.Context) {
	roleid := common.GetQueryToUint64(ctx, "roleid")
	menuIDList := []uint64{}
	where := sys.RoleMenu{RoleID: roleid}
	err := models.PluckList(&sys.RoleMenu{}, &where, &menuIDList, "menu_id")
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccess(ctx, &menuIDList)
}

// 设置角色菜单权限
func (Role) SetRole(ctx iris.Context) {
	roleid := common.GetQueryToUint64(ctx, "roleid")
	var menuids []uint64
	err := ctx.ReadJSON(&menuids)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	rm := sys.RoleMenu{}
	err = rm.SetRole(roleid, menuids)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	go casbins.CsbinSetRolePermission(roleid)
	common.ResSuccessMsg(ctx)
}
