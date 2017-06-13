package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ISOImageListCompleteArgs(ctx command.Context, params *params.ListISOImageParam, cur, prev, commandName string) {

}

func ISOImageCreateCompleteArgs(ctx command.Context, params *params.CreateISOImageParam, cur, prev, commandName string) {

}

func ISOImageReadCompleteArgs(ctx command.Context, params *params.ReadISOImageParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetCDROMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CDROMs {
		fmt.Println(res.CDROMs[i].ID)
		var target interface{} = &res.CDROMs[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ISOImageUpdateCompleteArgs(ctx command.Context, params *params.UpdateISOImageParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetCDROMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CDROMs {
		fmt.Println(res.CDROMs[i].ID)
		var target interface{} = &res.CDROMs[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ISOImageDeleteCompleteArgs(ctx command.Context, params *params.DeleteISOImageParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetCDROMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CDROMs {
		fmt.Println(res.CDROMs[i].ID)
		var target interface{} = &res.CDROMs[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ISOImageUploadCompleteArgs(ctx command.Context, params *params.UploadISOImageParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetCDROMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CDROMs {
		fmt.Println(res.CDROMs[i].ID)
		var target interface{} = &res.CDROMs[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ISOImageDownloadCompleteArgs(ctx command.Context, params *params.DownloadISOImageParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetCDROMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CDROMs {
		fmt.Println(res.CDROMs[i].ID)
		var target interface{} = &res.CDROMs[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ISOImageFtpOpenCompleteArgs(ctx command.Context, params *params.FtpOpenISOImageParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetCDROMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CDROMs {
		fmt.Println(res.CDROMs[i].ID)
		var target interface{} = &res.CDROMs[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ISOImageFtpCloseCompleteArgs(ctx command.Context, params *params.FtpCloseISOImageParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetCDROMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CDROMs {
		fmt.Println(res.CDROMs[i].ID)
		var target interface{} = &res.CDROMs[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}
