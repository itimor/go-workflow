package workflow

import (
	"go-workflow/backend/controllers/common"
	models "go-workflow/backend/models/common"
	"go-workflow/backend/models/workflow"

	"github.com/kataras/iris"
)

type CaseTypeStep struct{}

// 分页数据
func (CaseTypeStep) List(ctx iris.Context) {
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
	list := []workflow.CaseTypeStep{}
	err := models.GetPage(&workflow.CaseTypeStep{}, &workflow.CaseTypeStep{}, &list, page, limit, &total, whereOrder...)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccessPage(ctx, total, &list)
}

// 详情
func (CaseTypeStep) Detail(ctx iris.Context) {
	id := common.GetQueryToUint64(ctx, "id")
	var model workflow.CaseTypeStep
	where := workflow.CaseTypeStep{}
	where.ID = id
	_, err := models.First(&where, &model)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccess(ctx, &model)
}

// 更新
func (CaseTypeStep) Update(ctx iris.Context) {
	model := workflow.CaseTypeStep{}
	err := ctx.ReadJSON(&model)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	where := workflow.CaseTypeStep{}
	where.ID = model.ID
	modelOld := workflow.CaseTypeStep{}
	_, err = models.First(&where, &modelOld)
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
func (CaseTypeStep) Create(ctx iris.Context) {
	model := workflow.CaseTypeStep{}
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
	common.ResSuccess(ctx, model)
}

// 删除数据
func (CaseTypeStep) Delete(ctx iris.Context) {
	var ids, new_ids []uint64

	err := ctx.ReadJSON(&ids)
	if err != nil || len(ids) == 0 {
		common.ResErrSrv(ctx, err)
		return
	}

	obj := workflow.CaseTypeStep{}
	err = obj.Delete(new_ids)

	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccessMsg(ctx)
}

// 新增工作流类型后新建步骤
func InitStep(model workflow.CaseTypeStep) {
}

// 保存类型关联步骤
func (CaseTypeStep) CreateSteps(ctx iris.Context) {
	var steps []workflow.CaseTypeStep

	err := ctx.ReadJSON(&steps)
	if err != nil || len(steps) == 0 {
		common.ResErrSrv(ctx, err)
		return
	}

	for step, item := range steps {
		item.Step = step + 1
		println(item.CaseTypeID)
		err = models.Create(&item)
		if err != nil {
			common.ResFail(ctx, "操作失败")
			return
		}
	}
	common.ResSuccessMsg(ctx)
}
