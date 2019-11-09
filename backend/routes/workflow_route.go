package routes

import (
	"go-workflow/backend/controllers/workflow"

	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
)

func WorkflowRoute(party iris.Party) {
	api := party.Party("/api")
	{
		casetypes := workflow.CaseType{}
		api.PartyFunc("/casetype", func(casetype router.Party) {
			casetype.Get("/list", casetypes.List)
			casetype.Get("/detail", casetypes.Detail)
			casetype.Post("/delete", casetypes.Delete)
			casetype.Post("/update", casetypes.Update)
			casetype.Post("/create", casetypes.Create)
		})
		casetypesteps := workflow.CaseTypeStep{}
		api.PartyFunc("/casetypestep", func(casetypestep router.Party) {
			casetypestep.Get("/list", casetypesteps.List)
			casetypestep.Get("/detail", casetypesteps.Detail)
			casetypestep.Post("/delete", casetypesteps.Delete)
			casetypestep.Post("/update", casetypesteps.Update)
			casetypestep.Post("/create", casetypesteps.Create)
		})
		caseforms := workflow.CaseForm{}
		api.PartyFunc("/casetypestep", func(caseform router.Party) {
			caseform.Get("/list", caseforms.List)
			caseform.Get("/detail", caseforms.Detail)
			caseform.Post("/delete", caseforms.Delete)
			caseform.Post("/update", caseforms.Update)
			caseform.Post("/create", caseforms.Create)
		})
	}
}
