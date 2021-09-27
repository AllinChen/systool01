package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now().Format("20060102150405"))
		time.Sleep(time.Second)
		// if fmt.Sprint(time.Now().Clock()) == ""
	}

}
