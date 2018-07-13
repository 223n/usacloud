// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func DNSListCompleteArgs(ctx command.Context, params *params.ListDNSParam, cur, prev, commandName string) {

}

func DNSRecordInfoCompleteArgs(ctx command.Context, params *params.RecordInfoDNSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDNSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceDNSItems {
		fmt.Println(res.CommonServiceDNSItems[i].ID)
		var target interface{} = &res.CommonServiceDNSItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DNSCreateCompleteArgs(ctx command.Context, params *params.CreateDNSParam, cur, prev, commandName string) {

}

func DNSRecordAddCompleteArgs(ctx command.Context, params *params.RecordAddDNSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDNSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceDNSItems {
		fmt.Println(res.CommonServiceDNSItems[i].ID)
		var target interface{} = &res.CommonServiceDNSItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DNSReadCompleteArgs(ctx command.Context, params *params.ReadDNSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDNSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceDNSItems {
		fmt.Println(res.CommonServiceDNSItems[i].ID)
		var target interface{} = &res.CommonServiceDNSItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DNSRecordUpdateCompleteArgs(ctx command.Context, params *params.RecordUpdateDNSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDNSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceDNSItems {
		fmt.Println(res.CommonServiceDNSItems[i].ID)
		var target interface{} = &res.CommonServiceDNSItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DNSRecordDeleteCompleteArgs(ctx command.Context, params *params.RecordDeleteDNSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDNSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceDNSItems {
		fmt.Println(res.CommonServiceDNSItems[i].ID)
		var target interface{} = &res.CommonServiceDNSItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DNSUpdateCompleteArgs(ctx command.Context, params *params.UpdateDNSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDNSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceDNSItems {
		fmt.Println(res.CommonServiceDNSItems[i].ID)
		var target interface{} = &res.CommonServiceDNSItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DNSDeleteCompleteArgs(ctx command.Context, params *params.DeleteDNSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDNSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceDNSItems {
		fmt.Println(res.CommonServiceDNSItems[i].ID)
		var target interface{} = &res.CommonServiceDNSItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}
