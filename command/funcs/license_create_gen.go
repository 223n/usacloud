// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func LicenseCreate(ctx command.Context, params *params.CreateLicenseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLicenseAPI()
	p := api.New()

	// set params

	p.SetLicenseInfoByID(params.LicenseInfoId)

	p.SetName(params.Name)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("LicenseCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
