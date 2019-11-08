package workflowform

import (
	"fmt"

	"iris-ticket/backend/models/basemodel"
)

func TableName(name string) string {
	return fmt.Sprintf("%s%s%s", basemodel.GetTablePrefix(), "wform_", name)
}
