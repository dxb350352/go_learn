package main

import (
	"github.com/sas/utils"
	"github.com/sas/tqry/app/models"
	"fmt"
)

func main() {
	var inf models.Informer
	m := utils.GetObjectCT(inf,"colName")
	for k, v := range m {
		fmt.Println(k, "..", v)
	}
}

