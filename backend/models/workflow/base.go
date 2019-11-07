package workflow

import (
	"fmt"

	"iris-ticket/backend/models/basemodel"
)

func TableName(name string) string {
	return fmt.Sprintf("%s%s%s", basemodel.GetTablePrefix(), "wf_", name)
}
