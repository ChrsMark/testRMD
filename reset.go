package main

import (
	"github.com/intel/rmd/lib/resctrl"
	"fmt"
)

func main() {
	resGroups := resctrl.GetResAssociation()
	fmt.Println(resGroups)
	resctrl.DestroyResAssociation("COS3")
}
