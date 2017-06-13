// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProductInternetRead(ctx command.Context, params *params.ReadProductInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProductInternetAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProductInternetRead is failed: %s", e)
	}

	// set params

	// TODO Remove and implements here!!
	return fmt.Errorf("Not Implements ProductInternetRead : %#v", p)

	// call manipurate functions
	// res, err := api.XXXX(params.Id, p)
	// if err != nil {
	// 	return fmt.Errorf("ProductInternetRead is failed: %s", err)
	// }
	// return ctx.GetOutput().Print(res)

}
