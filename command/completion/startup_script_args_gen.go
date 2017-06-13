package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func StartupScriptListCompleteArgs(ctx command.Context, params *params.ListStartupScriptParam, cur, prev, commandName string) {

}

func StartupScriptCreateCompleteArgs(ctx command.Context, params *params.CreateStartupScriptParam, cur, prev, commandName string) {

}

func StartupScriptReadCompleteArgs(ctx command.Context, params *params.ReadStartupScriptParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetNoteAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Notes {
		fmt.Println(res.Notes[i].ID)
		var target interface{} = &res.Notes[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func StartupScriptUpdateCompleteArgs(ctx command.Context, params *params.UpdateStartupScriptParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetNoteAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Notes {
		fmt.Println(res.Notes[i].ID)
		var target interface{} = &res.Notes[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func StartupScriptDeleteCompleteArgs(ctx command.Context, params *params.DeleteStartupScriptParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetNoteAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Notes {
		fmt.Println(res.Notes[i].ID)
		var target interface{} = &res.Notes[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}
