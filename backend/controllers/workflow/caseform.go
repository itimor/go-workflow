package workflow

import (
	"go-workflow/backend/controllers/common"
	models "go-workflow/backend/models/common"
	"go-workflow/backend/models/workflow"

	"github.com/kataras/iris"
)

type CaseForm struct{}

// 分页数据
func (CaseForm) List(ctx iris.Context) {
	page := common.GetPageIndex(ctx)
	limit := common.GetPageLimit(ctx)
	sort := common.GetPageSort(ctx)
	key := common.GetPageKey(ctx)
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
	var total uint64
	list := []workflow.CaseForm{}
	err := models.GetPage(&workflow.CaseForm{}, &workflow.CaseForm{}, &list, page, limit, &total, whereOrder...)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccessPage(ctx, total, &list)
}

// 详情
func (CaseForm) Detail(ctx iris.Context) {
	id := common.GetQueryToUint64(ctx, "id")
	var model workflow.CaseForm
	where := workflow.CaseForm{}
	where.ID = id
	_, err := models.First(&where, &model)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccess(ctx, &model)
}

// 更新
func (CaseForm) Update(ctx iris.Context) {
	uid, _ := ctx.Values().GetUint64("auth_user_id")

	model := workflow.CaseForm{}
	err := ctx.ReadJSON(&model)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	where := workflow.CaseForm{}
	where.ID = model.ID
	modelOld := workflow.CaseForm{}
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
func (CaseForm) Create(ctx iris.Context) {
	uid, _ := ctx.Values().GetUint64("auth_user_id")

	model := workflow.CaseForm{}
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
func (CaseForm) Delete(ctx iris.Context) {
	var ids, new_ids []uint64

	err := ctx.ReadJSON(&ids)
	if err != nil || len(ids) == 0 {
		common.ResErrSrv(ctx, err)
		return
	}

	obj := workflow.CaseForm{}
	err = obj.Delete(new_ids)

	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccessMsg(ctx)
}
