package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func RegionListCompleteArgs(ctx command.Context, params *params.ListRegionParam, cur, prev, commandName string) {

}

func RegionReadCompleteArgs(ctx command.Context, params *params.ReadRegionParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	if cur != "" && !isSakuraID(cur) {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetRegionAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Regions {
		fmt.Println(res.Regions[i].ID)

	}

}
