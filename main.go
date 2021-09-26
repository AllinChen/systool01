package main

import (
	"fmt"

	"github.com/AllinChen/systool01/config"
)

func main() {
	fmt.Println(config.GlobalConf.Email)
}
