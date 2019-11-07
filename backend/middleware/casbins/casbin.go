package casbins

import (
	models "iris-ticket/backend/models/common"
	"iris-ticket/backend/models/sys"
	"iris-ticket/backend/pkg/convert"

	"github.com/casbin/casbin"
	casbinmodel "github.com/casbin/casbin/model"
)

const (
	PrefixUserID = "u"
	PrefixRoleID = "r"
)

var Enforcer *casbin.Enforcer

// 角色-URL导入
func InitCsbinEnforcer() (err error) {
	var enforcer *casbin.Enforcer
	casbinModel := `
	[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj, act
	
	[role_definition]
	g = _, _
	
	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.sub, p.sub) == true && keyMatch2(r.obj, p.obj) == true && regexMatch(r.act, p.act) == true || r.sub == "root"
	`
	m, _ := casbinmodel.NewModelFromString(casbinModel)
	enforcer, err = casbin.NewEnforcer(m)
	if err != nil {
		return
	}
	var roles []sys.Role
	err = models.Find(&sys.Role{}, &roles)
	if err != nil {
		return
	}
	if len(roles) == 0 {
		Enforcer = enforcer
		return
	}
	for _, role := range roles {
		setRolePermission(enforcer, role.ID)
	}
	Enforcer = enforcer
	return
}

// 删除角色
func CsbinDeleteRole(roleids []uint64) {
	if Enforcer == nil {
		return
	}
	for _, rid := range roleids {
		Enforcer.DeletePermissionsForUser(PrefixRoleID + convert.ToString(rid))
		Enforcer.DeleteRole(PrefixRoleID + convert.ToString(rid))
	}
}

// 设置角色权限
func CsbinSetRolePermission(roleid uint64) {
	if Enforcer == nil {
		return
	}
	Enforcer.DeletePermissionsForUser(PrefixRoleID + convert.ToString(roleid))
	setRolePermission(Enforcer, roleid)
}

// 设置角色权限
func setRolePermission(enforcer *casbin.Enforcer, roleid uint64) {
	var rolemenus []sys.RoleMenu
	err := models.Find(&sys.RoleMenu{RoleID: roleid}, &rolemenus)
	if err != nil {
		return
	}
	for _, rolemenu := range rolemenus {
		menu := sys.Menu{}
		where := sys.Menu{}
		where.ID = rolemenu.MenuID
		_, err = models.First(&where, &menu)
		if err != nil {
			return
		}
		if menu.MenuType == 3 {
			enforcer.AddPermissionForUser(PrefixRoleID+convert.ToString(roleid), "/api"+menu.URL, "GET|POST")
		}
	}
}

// 检查用户是否有权限
func CsbinCheckPermission(userID, url, methodtype string) (bool, error) {
	return Enforcer.Enforce(PrefixUserID+userID, url, methodtype)
}

// 用户角色处理
func CsbinAddRoleForUser(userid uint64) (err error) {
	if Enforcer == nil {
		return
	}
	uid := PrefixUserID + convert.ToString(userid)
	Enforcer.DeleteRolesForUser(uid)
	var userroles []sys.UserRole
	err = models.Find(&sys.UserRole{UserID: userid}, &userroles)
	if err != nil {
		return
	}
	for _, userrole := range userroles {
		Enforcer.AddRoleForUser(uid, PrefixRoleID+convert.ToString(userrole.RoleID))
	}
	return
}
