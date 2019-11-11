package routes

import (
	"go-workflow/backend/controllers/workflow"

	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
)

func WorkflowRoute(party iris.Party) {
	wfapi := party.Party("/workflow")
	{
		casetypes := workflow.CaseType{}
		wfapi.PartyFunc("/casetype", func(casetype router.Party) {
			casetype.Get("/list", casetypes.List)
			casetype.Get("/detail", casetypes.Detail)
			casetype.Post("/delete", casetypes.Delete)
			casetype.Post("/update", casetypes.Update)
			casetype.Post("/create", casetypes.Create)
		})
		casetypesteps := workflow.CaseTypeStep{}
		wfapi.PartyFunc("/casetypestep", func(casetypestep router.Party) {
			casetypestep.Get("/list", casetypesteps.List)
			casetypestep.Get("/detail", casetypesteps.Detail)
			casetypestep.Post("/delete", casetypesteps.Delete)
			casetypestep.Post("/update", casetypesteps.Update)
			casetypestep.Post("/create", casetypesteps.Create)
			casetypestep.Post("/createsteps", casetypesteps.CreateSteps)
		})
		caseforms := workflow.CaseForm{}
		wfapi.PartyFunc("/caseform", func(caseform router.Party) {
			caseform.Get("/list", caseforms.List)
			caseform.Get("/detail", caseforms.Detail)
			caseform.Post("/delete", caseforms.Delete)
			caseform.Post("/update", caseforms.Update)
			caseform.Post("/create", caseforms.Create)
		})
		cases := workflow.Case{}
		wfapi.PartyFunc("/case", func(wfcase router.Party) {
			wfcase.Get("/list", cases.List)
			wfcase.Get("/detail", cases.Detail)
			wfcase.Post("/delete", cases.Delete)
			wfcase.Post("/update", cases.Update)
			wfcase.Post("/create", cases.Create)
		})
		casesteps := workflow.CaseStep{}
		wfapi.PartyFunc("/casestep", func(casestep router.Party) {
			casestep.Get("/list", casesteps.List)
			casestep.Get("/detail", casesteps.Detail)
			casestep.Post("/delete", casesteps.Delete)
			casestep.Post("/update", casesteps.Update)
			casestep.Post("/create", casesteps.Create)
		})
		caseoperas := workflow.CaseOpera{}
		wfapi.PartyFunc("/caseopera", func(caseopera router.Party) {
			caseopera.Get("/list", caseoperas.List)
			caseopera.Get("/detail", caseoperas.Detail)
			caseopera.Post("/delete", caseoperas.Delete)
			caseopera.Post("/update", caseoperas.Update)
			caseopera.Post("/create", caseoperas.Create)
		})
	}
}
