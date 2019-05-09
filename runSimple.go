package main

import (
	"github.com/intel/rmd/lib/resctrl"
	"fmt"
)

func main() {
	resGroups := resctrl.GetResAssociation()
	fmt.Println(resGroups)
	rs := resctrl.NewResAssociation()
	rs.CPUs = "2"
	// mask = 00010000000

	group := []resctrl.CacheCos{{0, "80"}}
	rs.Schemata["COS2"] = group
	resctrl.Commit(rs, "COS2")
}
