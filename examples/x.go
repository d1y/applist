package main

import (
	"fmt"

	"github.com/d1y/applist"
)

func main() {
	var lists, _ = applist.GetApps()
	fmt.Println(lists)
}
