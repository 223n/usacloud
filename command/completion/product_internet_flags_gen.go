// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func ProductInternetListCompleteFlags(ctx command.Context, params *params.ListProductInternetParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		comp = define.Resources["ProductInternet"].Commands["list"].Params["name"].CompleteFunc
	case "id":
		comp = define.Resources["ProductInternet"].Commands["list"].Params["id"].CompleteFunc
	case "from", "offset":
		comp = define.Resources["ProductInternet"].Commands["list"].Params["from"].CompleteFunc
	case "max", "limit":
		comp = define.Resources["ProductInternet"].Commands["list"].Params["max"].CompleteFunc
	case "sort":
		comp = define.Resources["ProductInternet"].Commands["list"].Params["sort"].CompleteFunc
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProductInternetReadCompleteFlags(ctx command.Context, params *params.ReadProductInternetParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "assumeyes", "y":
		comp = define.Resources["ProductInternet"].Commands["read"].Params["assumeyes"].CompleteFunc
	case "id":
		comp = define.Resources["ProductInternet"].Commands["read"].Params["id"].CompleteFunc
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
