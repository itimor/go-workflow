package sys

import (
	"iris-ticket/backend/controllers/common"
	"iris-ticket/backend/middleware/casbins"
	models "iris-ticket/backend/models/common"
	"iris-ticket/backend/models/sys"
	"iris-ticket/backend/pkg/convert"

	"github.com/jameskeane/bcrypt"
	"github.com/kataras/iris"
)

type User struct{}

// 分页数据
func (User) List(ctx iris.Context) {
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
		whereOrder = append(whereOrder, models.PageWhereOrder{Where: "username like ? or realname like ?", Value: arr})
	}
	if status > 0 {
		var arr []interface{}
		arr = append(arr, status)
		whereOrder = append(whereOrder, models.PageWhereOrder{Where: "status = ?", Value: arr})
	}
	var total uint64
	list := []sys.User{}
	err := models.GetPage(&sys.User{}, &sys.User{}, &list, page, limit, &total, whereOrder...)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccessPage(ctx, total, &list)
}

// 详情
func (User) Detail(ctx iris.Context) {
	id := common.GetQueryToUint64(ctx, "id")
	var model sys.User
	where := sys.User{}
	where.ID = id
	_, err := models.First(&where, &model)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	model.Password = ""
	common.ResSuccess(ctx, &model)
}

// 更新
func (User) Update(ctx iris.Context) {
	model := sys.User{}
	err := ctx.ReadJSON(&model)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	where := sys.User{}
	where.ID = model.ID
	modelOld := sys.User{}
	_, err = models.First(&where, &modelOld)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	model.Username = modelOld.Username
	model.Password = modelOld.Password
	err = models.Save(&model)
	if err != nil {
		common.ResFail(ctx, "操作失败")
		return
	}
	common.ResSuccessMsg(ctx)
}

//新增
func (User) Create(ctx iris.Context) {
	model := sys.User{}
	err := ctx.ReadJSON(&model)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	salt, _ := bcrypt.Salt(10)
	hash, _ := bcrypt.Hash(model.Password, salt)
	model.Password = string(hash)
	err = models.Create(&model)
	if err != nil {
		common.ResFail(ctx, "操作失败")
		return
	}
	common.ResSuccess(ctx, model)
}

// 删除数据
func (User) Delete(ctx iris.Context) {
	var ids,new_ids []uint64

	err := ctx.ReadJSON(&ids)
	if err != nil || len(ids) == 0 {
		common.ResErrSrv(ctx, err)
		return
	}
	// 判断移除 super_admin id
	for _,i:= range ids {
		if common.SUPER_ADMIN_ID != convert.ToUint64(i) {
			new_ids = append(new_ids, convert.ToUint64(i))
		}
	}

	user := sys.User{}
	err = user.Delete(new_ids)

	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccessMsg(ctx)
}

// 获取用户下的角色ID列表
func (User) UserRoleIDList(ctx iris.Context) {
	UserID := common.GetQueryToUint64(ctx, "user_id")
	roleList := []uint64{}
	where := sys.UserRole{UserID: UserID}
	err := models.PluckList(&sys.UserRole{}, &where, &roleList, "role_id")
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccess(ctx, &roleList)
}

// 分配用户角色权限
func (User) SetRole(ctx iris.Context) {
	userid := common.GetQueryToUint64(ctx, "userid")
	var roleids []uint64
	err := ctx.ReadJSON(&roleids)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	ar := sys.UserRole{}
	err = ar.SetRole(userid, roleids)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	go casbins.CsbinAddRoleForUser(userid)
	common.ResSuccessMsg(ctx)
}
