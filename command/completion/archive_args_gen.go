package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ArchiveListCompleteArgs(ctx command.Context, params *params.ListArchiveParam, cur, prev, commandName string) {

}

func ArchiveCreateCompleteArgs(ctx command.Context, params *params.CreateArchiveParam, cur, prev, commandName string) {

}

func ArchiveReadCompleteArgs(ctx command.Context, params *params.ReadArchiveParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetArchiveAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Archives {
		fmt.Println(res.Archives[i].ID)
		var target interface{} = &res.Archives[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ArchiveUpdateCompleteArgs(ctx command.Context, params *params.UpdateArchiveParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetArchiveAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Archives {
		fmt.Println(res.Archives[i].ID)
		var target interface{} = &res.Archives[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ArchiveDeleteCompleteArgs(ctx command.Context, params *params.DeleteArchiveParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetArchiveAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Archives {
		fmt.Println(res.Archives[i].ID)
		var target interface{} = &res.Archives[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ArchiveUploadCompleteArgs(ctx command.Context, params *params.UploadArchiveParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetArchiveAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Archives {
		fmt.Println(res.Archives[i].ID)
		var target interface{} = &res.Archives[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ArchiveDownloadCompleteArgs(ctx command.Context, params *params.DownloadArchiveParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetArchiveAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Archives {
		fmt.Println(res.Archives[i].ID)
		var target interface{} = &res.Archives[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ArchiveFtpOpenCompleteArgs(ctx command.Context, params *params.FtpOpenArchiveParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetArchiveAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Archives {
		fmt.Println(res.Archives[i].ID)
		var target interface{} = &res.Archives[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ArchiveFtpCloseCompleteArgs(ctx command.Context, params *params.FtpCloseArchiveParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetArchiveAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Archives {
		fmt.Println(res.Archives[i].ID)
		var target interface{} = &res.Archives[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ArchiveWaitForCopyCompleteArgs(ctx command.Context, params *params.WaitForCopyArchiveParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetArchiveAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Archives {
		fmt.Println(res.Archives[i].ID)
		var target interface{} = &res.Archives[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}
