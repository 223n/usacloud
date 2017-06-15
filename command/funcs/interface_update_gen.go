// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func InterfaceUpdate(ctx command.Context, params *params.UpdateInterfaceParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInterfaceAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InterfaceUpdate is failed: %s", e)
	}

	// set params

	if ctx.IsSet("user-ipaddress") {
		p.SetUserIPAddress(params.UserIpaddress)
	}

	// call Update(id)
	res, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("InterfaceUpdate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}