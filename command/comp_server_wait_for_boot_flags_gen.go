// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func ServerWaitForBootCompleteFlags(ctx Context, params *WaitForBootServerParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["Server"].Commands["wait-for-boot"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}