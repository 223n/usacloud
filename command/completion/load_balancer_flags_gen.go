// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func LoadBalancerListCompleteFlags(ctx command.Context, params *params.ListLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		comp = define.Resources["LoadBalancer"].Commands["list"].Params["name"].CompleteFunc
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["list"].Params["id"].CompleteFunc
	case "tags":
		comp = define.Resources["LoadBalancer"].Commands["list"].Params["tags"].CompleteFunc
	case "from", "offset":
		comp = define.Resources["LoadBalancer"].Commands["list"].Params["from"].CompleteFunc
	case "max", "limit":
		comp = define.Resources["LoadBalancer"].Commands["list"].Params["max"].CompleteFunc
	case "sort":
		comp = define.Resources["LoadBalancer"].Commands["list"].Params["sort"].CompleteFunc
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

func LoadBalancerCreateCompleteFlags(ctx command.Context, params *params.CreateLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "switch-id":
		comp = define.Resources["LoadBalancer"].Commands["create"].Params["switch-id"].CompleteFunc
	case "vrid", "VRID":
		comp = define.Resources["LoadBalancer"].Commands["create"].Params["vrid"].CompleteFunc
	case "high-availability", "ha":
		comp = define.Resources["LoadBalancer"].Commands["create"].Params["high-availability"].CompleteFunc
	case "plan":
		comp = define.Resources["LoadBalancer"].Commands["create"].Params["plan"].CompleteFunc
	case "ipaddress1", "ip1":
		comp = define.Resources["LoadBalancer"].Commands["create"].Params["ipaddress1"].CompleteFunc
	case "ipaddress2", "ip2":
		comp = define.Resources["LoadBalancer"].Commands["create"].Params["ipaddress2"].CompleteFunc
	case "nw-mask-len":
		comp = define.Resources["LoadBalancer"].Commands["create"].Params["nw-mask-len"].CompleteFunc
	case "default-route":
		comp = define.Resources["LoadBalancer"].Commands["create"].Params["default-route"].CompleteFunc
	case "name":
		comp = define.Resources["LoadBalancer"].Commands["create"].Params["name"].CompleteFunc
	case "description", "desc":
		comp = define.Resources["LoadBalancer"].Commands["create"].Params["description"].CompleteFunc
	case "tags":
		comp = define.Resources["LoadBalancer"].Commands["create"].Params["tags"].CompleteFunc
	case "icon-id":
		comp = define.Resources["LoadBalancer"].Commands["create"].Params["icon-id"].CompleteFunc
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

func LoadBalancerReadCompleteFlags(ctx command.Context, params *params.ReadLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["read"].Params["id"].CompleteFunc
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

func LoadBalancerUpdateCompleteFlags(ctx command.Context, params *params.UpdateLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		comp = define.Resources["LoadBalancer"].Commands["update"].Params["name"].CompleteFunc
	case "description", "desc":
		comp = define.Resources["LoadBalancer"].Commands["update"].Params["description"].CompleteFunc
	case "tags":
		comp = define.Resources["LoadBalancer"].Commands["update"].Params["tags"].CompleteFunc
	case "icon-id":
		comp = define.Resources["LoadBalancer"].Commands["update"].Params["icon-id"].CompleteFunc
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["update"].Params["id"].CompleteFunc
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

func LoadBalancerDeleteCompleteFlags(ctx command.Context, params *params.DeleteLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "force", "f":
		comp = define.Resources["LoadBalancer"].Commands["delete"].Params["force"].CompleteFunc
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["delete"].Params["id"].CompleteFunc
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

func LoadBalancerBootCompleteFlags(ctx command.Context, params *params.BootLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["boot"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func LoadBalancerShutdownCompleteFlags(ctx command.Context, params *params.ShutdownLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["shutdown"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func LoadBalancerShutdownForceCompleteFlags(ctx command.Context, params *params.ShutdownForceLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["shutdown-force"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func LoadBalancerResetCompleteFlags(ctx command.Context, params *params.ResetLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["reset"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func LoadBalancerWaitForBootCompleteFlags(ctx command.Context, params *params.WaitForBootLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["wait-for-boot"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func LoadBalancerWaitForDownCompleteFlags(ctx command.Context, params *params.WaitForDownLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["wait-for-down"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func LoadBalancerVipInfoCompleteFlags(ctx command.Context, params *params.VipInfoLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["vip-info"].Params["id"].CompleteFunc
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

func LoadBalancerVipAddCompleteFlags(ctx command.Context, params *params.VipAddLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "vip":
		comp = define.Resources["LoadBalancer"].Commands["vip-add"].Params["vip"].CompleteFunc
	case "port":
		comp = define.Resources["LoadBalancer"].Commands["vip-add"].Params["port"].CompleteFunc
	case "delay-loop":
		comp = define.Resources["LoadBalancer"].Commands["vip-add"].Params["delay-loop"].CompleteFunc
	case "sorry-server":
		comp = define.Resources["LoadBalancer"].Commands["vip-add"].Params["sorry-server"].CompleteFunc
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["vip-add"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func LoadBalancerVipUpdateCompleteFlags(ctx command.Context, params *params.VipUpdateLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		comp = define.Resources["LoadBalancer"].Commands["vip-update"].Params["index"].CompleteFunc
	case "vip":
		comp = define.Resources["LoadBalancer"].Commands["vip-update"].Params["vip"].CompleteFunc
	case "port":
		comp = define.Resources["LoadBalancer"].Commands["vip-update"].Params["port"].CompleteFunc
	case "delay-loop":
		comp = define.Resources["LoadBalancer"].Commands["vip-update"].Params["delay-loop"].CompleteFunc
	case "sorry-server":
		comp = define.Resources["LoadBalancer"].Commands["vip-update"].Params["sorry-server"].CompleteFunc
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["vip-update"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func LoadBalancerVipDeleteCompleteFlags(ctx command.Context, params *params.VipDeleteLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		comp = define.Resources["LoadBalancer"].Commands["vip-delete"].Params["index"].CompleteFunc
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["vip-delete"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func LoadBalancerServerInfoCompleteFlags(ctx command.Context, params *params.ServerInfoLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "vip-index":
		comp = define.Resources["LoadBalancer"].Commands["server-info"].Params["vip-index"].CompleteFunc
	case "vip":
		comp = define.Resources["LoadBalancer"].Commands["server-info"].Params["vip"].CompleteFunc
	case "port":
		comp = define.Resources["LoadBalancer"].Commands["server-info"].Params["port"].CompleteFunc
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["server-info"].Params["id"].CompleteFunc
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

func LoadBalancerServerAddCompleteFlags(ctx command.Context, params *params.ServerAddLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "vip-index":
		comp = define.Resources["LoadBalancer"].Commands["server-add"].Params["vip-index"].CompleteFunc
	case "vip":
		comp = define.Resources["LoadBalancer"].Commands["server-add"].Params["vip"].CompleteFunc
	case "port":
		comp = define.Resources["LoadBalancer"].Commands["server-add"].Params["port"].CompleteFunc
	case "ipaddress", "ip":
		comp = define.Resources["LoadBalancer"].Commands["server-add"].Params["ipaddress"].CompleteFunc
	case "protocol":
		comp = define.Resources["LoadBalancer"].Commands["server-add"].Params["protocol"].CompleteFunc
	case "path":
		comp = define.Resources["LoadBalancer"].Commands["server-add"].Params["path"].CompleteFunc
	case "response-code":
		comp = define.Resources["LoadBalancer"].Commands["server-add"].Params["response-code"].CompleteFunc
	case "enabled":
		comp = define.Resources["LoadBalancer"].Commands["server-add"].Params["enabled"].CompleteFunc
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["server-add"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func LoadBalancerServerUpdateCompleteFlags(ctx command.Context, params *params.ServerUpdateLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "vip-index":
		comp = define.Resources["LoadBalancer"].Commands["server-update"].Params["vip-index"].CompleteFunc
	case "vip":
		comp = define.Resources["LoadBalancer"].Commands["server-update"].Params["vip"].CompleteFunc
	case "port":
		comp = define.Resources["LoadBalancer"].Commands["server-update"].Params["port"].CompleteFunc
	case "ipaddress", "ip":
		comp = define.Resources["LoadBalancer"].Commands["server-update"].Params["ipaddress"].CompleteFunc
	case "protocol":
		comp = define.Resources["LoadBalancer"].Commands["server-update"].Params["protocol"].CompleteFunc
	case "path":
		comp = define.Resources["LoadBalancer"].Commands["server-update"].Params["path"].CompleteFunc
	case "response-code":
		comp = define.Resources["LoadBalancer"].Commands["server-update"].Params["response-code"].CompleteFunc
	case "enabled":
		comp = define.Resources["LoadBalancer"].Commands["server-update"].Params["enabled"].CompleteFunc
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["server-update"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func LoadBalancerServerDeleteCompleteFlags(ctx command.Context, params *params.ServerDeleteLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "vip-index":
		comp = define.Resources["LoadBalancer"].Commands["server-delete"].Params["vip-index"].CompleteFunc
	case "vip":
		comp = define.Resources["LoadBalancer"].Commands["server-delete"].Params["vip"].CompleteFunc
	case "port":
		comp = define.Resources["LoadBalancer"].Commands["server-delete"].Params["port"].CompleteFunc
	case "ipaddress", "ip":
		comp = define.Resources["LoadBalancer"].Commands["server-delete"].Params["ipaddress"].CompleteFunc
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["server-delete"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func LoadBalancerMonitorCompleteFlags(ctx command.Context, params *params.MonitorLoadBalancerParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "start":
		comp = define.Resources["LoadBalancer"].Commands["monitor"].Params["start"].CompleteFunc
	case "end":
		comp = define.Resources["LoadBalancer"].Commands["monitor"].Params["end"].CompleteFunc
	case "key-format":
		comp = define.Resources["LoadBalancer"].Commands["monitor"].Params["key-format"].CompleteFunc
	case "id":
		comp = define.Resources["LoadBalancer"].Commands["monitor"].Params["id"].CompleteFunc
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
