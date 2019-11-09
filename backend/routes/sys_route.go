package routes

import (
	"go-workflow/backend/controllers/sys"

	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
)

func SysRoute(party iris.Party) {
	sysapi := party.Party("/sys")
	{
		auths := sys.Auth{}
		sysapi.PartyFunc("/auth", func(auth router.Party) {
			auth.Get("/info", auths.Info)
			auth.Post("/login", auths.Login)
			auth.Post("/logout", auths.Logout)
			auth.Post("/changepwd", auths.ChangePwd)
		})
		menus := sys.Menu{}
		sysapi.PartyFunc("/menu", func(menu router.Party) {
			menu.Get("/list", menus.List)
			menu.Get("/detail", menus.Detail)
			menu.Get("/allmenu", menus.AllMenu)
			menu.Get("/menubuttonlist", menus.MenuButtonList)
			menu.Post("/delete", menus.Delete)
			menu.Post("/update", menus.Update)
			menu.Post("/create", menus.Create)
		})
		users := sys.User{}
		sysapi.PartyFunc("/user", func(user router.Party) {
			user.Get("/detail", users.Detail)
			user.Get("/list", users.List)
			user.Get("/userroleidlist", users.UserRoleIDList)
			user.Post("/delete", users.Delete)
			user.Post("/update", users.Update)
			user.Post("/create", users.Create)
			user.Post("/setrole", users.SetRole)
		})
		roles := sys.Role{}
		sysapi.PartyFunc("/role", func(role router.Party) {
			role.Get("/list", roles.List)
			role.Get("/detail", roles.Detail)
			role.Get("/rolemenuidlist", roles.RoleMenuIDList)
			role.Get("/allrole", roles.AllRole)
			role.Post("/delete", roles.Delete)
			role.Post("/update", roles.Update)
			role.Post("/create", roles.Create)
			role.Post("/setrole", roles.SetRole)
		})
	}
}
