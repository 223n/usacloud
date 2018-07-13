// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SIMDelete(ctx command.Context, params *params.DeleteSIMParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSIMAPI()

	// set params

	// call Delete(id)
	res, err := api.Delete(params.Id)
	if err != nil {
		return fmt.Errorf("SIMDelete is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
