package workflow

import (
	"go-workflow/backend/controllers/common"
	models "go-workflow/backend/models/common"
	"go-workflow/backend/models/workflow"

	"github.com/kataras/iris"
)

type CaseOpera struct{}

// 分页数据
func (CaseOpera) List(ctx iris.Context) {
	page := common.GetPageIndex(ctx)
	limit := common.GetPageLimit(ctx)
	sort := common.GetPageSort(ctx)
	key := common.GetPageKey(ctx)
	status := common.GetQueryToUint(ctx, "status")
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
		whereOrder = append(whereOrder, models.PageWhereOrder{Where: "name like ?", Value: arr})
	}
	if status > 0 {
		var arr []interface{}
		arr = append(arr, status)
		whereOrder = append(whereOrder, models.PageWhereOrder{Where: "status = ?", Value: arr})
	}
	var total uint64
	list := []workflow.CaseOpera{}
	err := models.GetPage(&workflow.CaseOpera{}, &workflow.CaseOpera{}, &list, page, limit, &total, whereOrder...)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccessPage(ctx, total, &list)
}

// 详情
func (CaseOpera) Detail(ctx iris.Context) {
	id := common.GetQueryToUint64(ctx, "id")
	var model workflow.CaseOpera
	where := workflow.CaseOpera{}
	where.ID = id
	_, err := models.First(&where, &model)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccess(ctx, &model)
}

// 更新
func (CaseOpera) Update(ctx iris.Context) {
	uid, _ := ctx.Values().GetUint64("auth_user_id")

	model := workflow.CaseOpera{}
	err := ctx.ReadJSON(&model)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	where := workflow.CaseOpera{}
	where.ID = model.ID
	modelOld := workflow.CaseOpera{}
	_, err = models.First(&where, &modelOld)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	model.UpdatedBy = uid
	err = models.Save(&model)
	if err != nil {
		common.ResFail(ctx, "操作失败")
		return
	}
	common.ResSuccessMsg(ctx)
}

//新增
func (CaseOpera) Create(ctx iris.Context) {
	uid, _ := ctx.Values().GetUint64("auth_user_id")

	model := workflow.CaseOpera{}
	err := ctx.ReadJSON(&model)

	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	model.CreatedBy = uid
	err = models.Create(&model)
	if err != nil {
		common.ResFail(ctx, "操作失败")
		return
	}
	common.ResSuccess(ctx, model)
}

// 删除数据
func (CaseOpera) Delete(ctx iris.Context) {
	var ids, new_ids []uint64

	err := ctx.ReadJSON(&ids)
	if err != nil || len(ids) == 0 {
		common.ResErrSrv(ctx, err)
		return
	}

	obj := workflow.CaseOpera{}
	err = obj.Delete(new_ids)

	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccessMsg(ctx)
}
