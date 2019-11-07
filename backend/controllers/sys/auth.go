package sys

import (
	"time"

	"iris-ticket/backend/config"
	"iris-ticket/backend/controllers/common"
	models "iris-ticket/backend/models/common"
	"iris-ticket/backend/models/sys"
	"iris-ticket/backend/pkg/convert"

	"github.com/ahmetb/go-linq"
	"github.com/dgrijalva/jwt-go"
	"github.com/jameskeane/bcrypt"
	"github.com/kataras/iris"
)

type Auth struct{}

// 用户登录
func (Auth) Login(ctx iris.Context) {
	aul := sys.User{}
	if err := ctx.ReadJSON(&aul); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		common.ResErrSrv(ctx, err)
		return
	} else {
		if UserNameErr := common.Validate.Var(aul.Username, "required,min=3,max=20"); UserNameErr != nil {
			common.ResFail(ctx, "username format err")
			return
		} else if PwdErr := common.Validate.Var(aul.Password, "required,min=6,max=20"); PwdErr != nil {
			common.ResFail(ctx, "password format err")
		} else {
			ctx.StatusCode(iris.StatusOK)
			if response, status, _ := CheckLogin(ctx, aul.Username, aul.Password); status {
				common.ResSuccess(ctx, response)
			}
		}
	}
}

// 用户登出
func (Auth) Logout(ctx iris.Context) {
	// 删除uid
	uid, _ := ctx.Values().GetUint64("auth_user_id")
	where := sys.OauthToken{}
	where.Revoked = false
	where.UserId = uid
	modelOld := sys.OauthToken{}
	_, err := models.First(&where, &modelOld)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	modelNew := sys.OauthToken{Revoked: true}
	err = models.Updates(&modelOld, &modelNew)
	if err != nil {
		common.ResErrSrv(ctx, err)
		return
	}
	common.ResSuccess(ctx, uid)
}

// 用户修改密码
func (Auth) ChangePwd(ctx iris.Context) {
	modelNew := sys.User{}
	if err := ctx.ReadJSON(&modelNew); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		common.ResErrSrv(ctx, err)
		return
	} else {
		uid, _ := ctx.Values().GetUint64("auth_user_id")
		where := sys.User{}
		where.ID = uid
		modelOld := sys.User{}
		_, err := models.First(&where, &modelOld)
		if err != nil {
			common.ResErrSrv(ctx, err)
			return
		}
		salt, _ := bcrypt.Salt(10)
		hash, _ := bcrypt.Hash(modelNew.Password, salt)
		modelNew.Password = string(hash)
		err = models.Updates(&modelOld, &modelNew)
		if err != nil {
			common.ResErrSrv(ctx, err)
			return
		}
		common.ResSuccess(ctx, "password change success")
	}
}

/**
 * 校验用户登录
 * @method UserAdminCheckLogin
 * @param  {[type]}  username string [description]
 */
func UserAdminCheckLogin(ctx iris.Context, username string) (model sys.User) {
	where := sys.User{}
	where.Username = username
	model = sys.User{}
	_, err := models.First(&where, &model)
	if err != nil {
		common.ResFail(ctx, "操作失败")
		return
	}
	return
}

/**
 * 判断用户是否登录
 * @method CheckLogin
 * @param  {[type]}  username string    [description]
 * @param  {[type]}  password string [description]
 */
func CheckLogin(ctx iris.Context, username, password string) (response string, status bool, msg string) {
	user := UserAdminCheckLogin(ctx, username)
	if user.ID == 0 {
		status = false
		msg = "user is not exist"
		return
	} else {
		if ok := bcrypt.Match(password, user.Password); ok {
			expireTime := time.Now().Add(time.Hour * time.Duration(config.Conf.Get("jwt.timeout").(int64))).Unix()
			jwtSecret := config.Conf.Get("jwt.secert").(string)
			token := jwt.New(jwt.SigningMethodHS256)
			claims := make(jwt.MapClaims)
			claims["exp"] = expireTime
			claims["iat"] = time.Now().Unix()
			token.Claims = claims
			Tokenstring, err := token.SignedString([]byte(jwtSecret))

			if err != nil {
				common.ResFail(ctx, err.Error())
				return
			}

			oauthToken := new(sys.OauthToken)
			oauthToken.Token = Tokenstring
			oauthToken.UserId = user.ID
			oauthToken.Secret = jwtSecret
			oauthToken.Revoked = false
			oauthToken.ExpressIn = expireTime
			oauthToken.CreatedAt = time.Now()
			err = models.Create(&oauthToken)
			if err != nil {
				status = false
			} else {
				response = Tokenstring
				status = true
			}
			return
		} else {
			common.ResFail(ctx, "密码错误")
			return
		}
	}
}

type MenuMeta struct {
	Title   string `json:"title"`   // 标题
	Icon    string `json:"icon"`    // 图标
	NoCache bool   `json:"noCache"` // 是不是缓存
}

type MenuModel struct {
	Path      string      `json:"path"`      // 路由
	Component string      `json:"component"` // 对应vue中的map name
	Name      string      `json:"name"`      // 菜单名称
	Hidden    bool        `json:"hidden"`    // 是否隐藏
	Meta      MenuMeta    `json:"meta"`      // 菜单信息
	Children  []MenuModel `json:"children"`  // 子级菜单
}

type UserData struct {
	Menus        []MenuModel `json:"menus"`        // 菜单
	Introduction string      `json:"introduction"` // 介绍
	Avatar       string      `json:"avatar"`       // 图标
	Username     string      `json:"username"`     // 姓名
}

// 获取用户信息及可访问的权限菜单
func (Auth) Info(ctx iris.Context) {
	// 用户ID
	uid, err := ctx.Values().GetUint64("auth_user_id")
	if err != nil {
		common.ResFailCode(ctx, "token 无效", 50008)
		return
	}
	userID := convert.ToUint64(uid)
	where := sys.User{}
	where.ID = userID
	model := sys.User{}
	models.First(&where, &model)
	// 根据用户ID获取用户权限菜单
	var menuData []sys.Menu
	if userID == common.SUPER_ADMIN_ID {
		//管理员
		menuData, err = getAllMenu()
		if err != nil {
			common.ResErrSrv(ctx, err)
			return
		}
		if len(menuData) == 0 {
			menuModelTop := sys.Menu{Status: 1, ParentID: 0, URL: "", Name: "TOP", Sequence: 1, MenuType: 1, Code: "TOP", OperateType: "none"}
			models.Create(&menuModelTop)
			menuModelSys := sys.Menu{Status: 1, ParentID: menuModelTop.ID, URL: "", Name: "系统管理", Sequence: 1, MenuType: 1, Code: "Sys", Icon: "lock", OperateType: "none"}
			models.Create(&menuModelSys)
			menuModel := sys.Menu{Status: 1, ParentID: menuModelSys.ID, URL: "/menu", Name: "菜单管理", Sequence: 20, MenuType: 2, Code: "Menu", Icon: "documentation", OperateType: "none"}
			models.Create(&menuModel)
			InitMenu(menuModel)
			menuModel = sys.Menu{Status: 1, ParentID: menuModelSys.ID, URL: "/role", Name: "角色管理", Sequence: 30, MenuType: 2, Code: "Role", Icon: "tree", OperateType: "none"}
			models.Create(&menuModel)
			InitMenu(menuModel)
			menuModel = sys.Menu{Status: 1, ParentID: menuModel.ID, URL: "/role/setrole", Name: "分配角色菜单", Sequence: 6, MenuType: 3, Code: "RoleSetrolemenu", Icon: "", OperateType: "setrolemenu"}
			models.Create(&menuModel)
			menuModel = sys.Menu{Status: 1, ParentID: menuModelSys.ID, URL: "/user", Name: "用户管理", Sequence: 40, MenuType: 2, Code: "user", Icon: "user", OperateType: "none"}
			models.Create(&menuModel)
			InitMenu(menuModel)
			menuModel = sys.Menu{Status: 1, ParentID: menuModel.ID, URL: "/user/setrole", Name: "分配角色", Sequence: 6, MenuType: 3, Code: "userSetrole", Icon: "", OperateType: "setadminrole"}
			models.Create(&menuModel)

			menuData, _ = getAllMenu()
		}
	} else {
		menuData, err = getMenusByUserid(userID)
		if err != nil {
			common.ResErrSrv(ctx, err)
			return
		}
	}
	var menus []MenuModel
	if len(menuData) > 0 {
		var topmenuid uint64 = menuData[0].ParentID
		if topmenuid == 0 {
			topmenuid = menuData[0].ID
		}
		menus = setMenu(menuData, topmenuid)
	}
	if len(menus) == 0 && userID == common.SUPER_ADMIN_ID {
		menus = getSuperAdminMenu()
	}
	resData := UserData{Menus: menus, Username: model.Realname, Avatar: model.Avatar}
	common.ResSuccess(ctx, &resData)
}

//查询所有菜单
func getAllMenu() (menus []sys.Menu, err error) {
	models.Find(&sys.Menu{}, &menus, "parent_id asc", "sequence asc")
	return
}

//获取超级管理员初使菜单
func getSuperAdminMenu() (out []MenuModel) {
	menuTop := MenuModel{
		Path:      "/sys",
		Component: "Sys",
		Name:      "Sys",
		Meta:      MenuMeta{Title: "系统管理", NoCache: false},
		Children:  []MenuModel{}}
	menuModel := MenuModel{
		Path:      "/menu",
		Component: "Menu",
		Name:      "Menu",
		Meta:      MenuMeta{Title: "菜单管理", NoCache: false},
		Children:  []MenuModel{}}
	menuTop.Children = append(menuTop.Children, menuModel)
	menuModel = MenuModel{
		Path:      "/role",
		Component: "Role",
		Name:      "Role",
		Meta:      MenuMeta{Title: "角色管理", NoCache: false},
		Children:  []MenuModel{}}
	menuTop.Children = append(menuTop.Children, menuModel)
	menuModel = MenuModel{
		Path:      "/user",
		Component: "user",
		Name:      "user",
		Meta:      MenuMeta{Title: "用户管理", NoCache: false},
		Children:  []MenuModel{}}
	menuTop.Children = append(menuTop.Children, menuModel)
	out = append(out, menuTop)
	return
}

// 递归菜单
func setMenu(menus []sys.Menu, parentID uint64) (out []MenuModel) {
	var menuArr []sys.Menu
	linq.From(menus).Where(func(ctx interface{}) bool {
		return ctx.(sys.Menu).ParentID == parentID
	}).OrderBy(func(ctx interface{}) interface{} {
		return ctx.(sys.Menu).Sequence
	}).ToSlice(&menuArr)
	if len(menuArr) == 0 {
		return
	}
	noCache := false
	for _, item := range menuArr {
		menu := MenuModel{
			Path:      item.URL,
			Component: item.Code,
			Name:      item.Code,
			Meta:      MenuMeta{Title: item.Name, Icon: item.Icon, NoCache: noCache},
			Children:  []MenuModel{}}
		if item.MenuType == 3 {
			menu.Hidden = true
		}
		//查询是否有子级
		menuChildren := setMenu(menus, item.ID)
		if len(menuChildren) > 0 {
			menu.Children = menuChildren
		}
		if item.MenuType == 2 {
			// 添加子级首页，有这一级NoCache才有效
			menuIndex := MenuModel{
				Path:      "index",
				Component: item.Code,
				Name:      item.Code,
				Meta:      MenuMeta{Title: item.Name, Icon: item.Icon, NoCache: noCache},
				Children:  []MenuModel{}}
			menu.Children = append(menu.Children, menuIndex)
			menu.Name = menu.Name + "index"
			menu.Meta = MenuMeta{}
		}
		out = append(out, menu)
	}
	return
}

//查询登录用户权限菜单
func getMenusByUserid(userid uint64) (ret []sys.Menu, err error) {
	menu := sys.Menu{}
	var menus []sys.Menu
	err = menu.GetMenuByUserid(userid, &menus)
	if err != nil || len(menus) == 0 {
		return
	}
	allmenu, err := getAllMenu()
	if err != nil || len(allmenu) == 0 {
		return
	}
	menuMapAll := make(map[uint64]sys.Menu)
	for _, item := range allmenu {
		menuMapAll[item.ID] = item
	}
	menuMap := make(map[uint64]sys.Menu)
	for _, item := range menus {
		menuMap[item.ID] = item
	}
	for _, item := range menus {
		_, exists := menuMap[item.ParentID]
		if exists {
			continue
		}
		setMenuUp(menuMapAll, item.ParentID, menuMap)
	}
	for _, m := range menuMap {
		ret = append(ret, m)
	}
	linq.From(ret).OrderBy(func(ctx interface{}) interface{} {
		return ctx.(sys.Menu).ParentID
	}).ToSlice(&ret)
	return
}

// 向上查找父级菜单
func setMenuUp(menuMapAll map[uint64]sys.Menu, menuid uint64, menuMap map[uint64]sys.Menu) {
	menuModel, exists := menuMapAll[menuid]
	if exists {
		mid := menuModel.ID
		_, exists = menuMap[mid]
		if !exists {
			menuMap[mid] = menuModel
			setMenuUp(menuMapAll, menuModel.ParentID, menuMap)
		}
	}
}
