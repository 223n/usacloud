package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func PacketFilterListCompleteArgs(ctx command.Context, params *params.ListPacketFilterParam, cur, prev, commandName string) {

}

func PacketFilterCreateCompleteArgs(ctx command.Context, params *params.CreatePacketFilterParam, cur, prev, commandName string) {

}

func PacketFilterReadCompleteArgs(ctx command.Context, params *params.ReadPacketFilterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPacketFilterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.PacketFilters {
		fmt.Println(res.PacketFilters[i].ID)
		var target interface{} = &res.PacketFilters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func PacketFilterUpdateCompleteArgs(ctx command.Context, params *params.UpdatePacketFilterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPacketFilterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.PacketFilters {
		fmt.Println(res.PacketFilters[i].ID)
		var target interface{} = &res.PacketFilters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func PacketFilterDeleteCompleteArgs(ctx command.Context, params *params.DeletePacketFilterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPacketFilterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.PacketFilters {
		fmt.Println(res.PacketFilters[i].ID)
		var target interface{} = &res.PacketFilters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func PacketFilterRuleInfoCompleteArgs(ctx command.Context, params *params.RuleInfoPacketFilterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPacketFilterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.PacketFilters {
		fmt.Println(res.PacketFilters[i].ID)
		var target interface{} = &res.PacketFilters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func PacketFilterRuleAddCompleteArgs(ctx command.Context, params *params.RuleAddPacketFilterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPacketFilterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.PacketFilters {
		fmt.Println(res.PacketFilters[i].ID)
		var target interface{} = &res.PacketFilters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func PacketFilterRuleUpdateCompleteArgs(ctx command.Context, params *params.RuleUpdatePacketFilterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPacketFilterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.PacketFilters {
		fmt.Println(res.PacketFilters[i].ID)
		var target interface{} = &res.PacketFilters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func PacketFilterRuleDeleteCompleteArgs(ctx command.Context, params *params.RuleDeletePacketFilterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPacketFilterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.PacketFilters {
		fmt.Println(res.PacketFilters[i].ID)
		var target interface{} = &res.PacketFilters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func PacketFilterInterfaceConnectCompleteArgs(ctx command.Context, params *params.InterfaceConnectPacketFilterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPacketFilterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.PacketFilters {
		fmt.Println(res.PacketFilters[i].ID)
		var target interface{} = &res.PacketFilters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func PacketFilterInterfaceDisconnectCompleteArgs(ctx command.Context, params *params.InterfaceDisconnectPacketFilterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPacketFilterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.PacketFilters {
		fmt.Println(res.PacketFilters[i].ID)
		var target interface{} = &res.PacketFilters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}
