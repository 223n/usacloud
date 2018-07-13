// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func InterfaceListCompleteFlags(ctx command.Context, params *params.ListInterfaceParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["Interface"].Commands["list"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Interface"].Commands["list"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "from", "offset":
		param := define.Resources["Interface"].Commands["list"].BuildedParams().Get("from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "max", "limit":
		param := define.Resources["Interface"].Commands["list"].BuildedParams().Get("max")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sort":
		param := define.Resources["Interface"].Commands["list"].BuildedParams().Get("sort")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
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

func InterfacePacketFilterConnectCompleteFlags(ctx command.Context, params *params.PacketFilterConnectInterfaceParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "packet-filter-id":
		param := define.Resources["Interface"].Commands["packet-filter-connect"].BuildedParams().Get("packet-filter-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Interface"].Commands["packet-filter-connect"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func InterfaceCreateCompleteFlags(ctx command.Context, params *params.CreateInterfaceParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "server-id":
		param := define.Resources["Interface"].Commands["create"].BuildedParams().Get("server-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
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

func InterfacePacketFilterDisconnectCompleteFlags(ctx command.Context, params *params.PacketFilterDisconnectInterfaceParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "packet-filter-id":
		param := define.Resources["Interface"].Commands["packet-filter-disconnect"].BuildedParams().Get("packet-filter-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Interface"].Commands["packet-filter-disconnect"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func InterfaceReadCompleteFlags(ctx command.Context, params *params.ReadInterfaceParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		param := define.Resources["Interface"].Commands["read"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
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

func InterfaceUpdateCompleteFlags(ctx command.Context, params *params.UpdateInterfaceParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "user-ipaddress":
		param := define.Resources["Interface"].Commands["update"].BuildedParams().Get("user-ipaddress")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Interface"].Commands["update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
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

func InterfaceDeleteCompleteFlags(ctx command.Context, params *params.DeleteInterfaceParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		param := define.Resources["Interface"].Commands["delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
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
