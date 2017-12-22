package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func WebAccelListCompleteArgs(ctx command.Context, params *params.ListWebAccelParam, cur, prev, commandName string) {

}

func WebAccelReadCompleteArgs(ctx command.Context, params *params.ReadWebAccelParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetWebAccelAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.WebAccelSites {
		fmt.Println(res.WebAccelSites[i].ID)
		var target interface{} = &res.WebAccelSites[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func WebAccelCertificateInfoCompleteArgs(ctx command.Context, params *params.CertificateInfoWebAccelParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetWebAccelAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.WebAccelSites {
		fmt.Println(res.WebAccelSites[i].ID)
		var target interface{} = &res.WebAccelSites[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func WebAccelCertificateUpdateCompleteArgs(ctx command.Context, params *params.CertificateUpdateWebAccelParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetWebAccelAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.WebAccelSites {
		fmt.Println(res.WebAccelSites[i].ID)
		var target interface{} = &res.WebAccelSites[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func WebAccelDeleteCacheCompleteArgs(ctx command.Context, params *params.DeleteCacheWebAccelParam, cur, prev, commandName string) {

}
