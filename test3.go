package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(TimeConvert("12:01:00PM"))
	fmt.Println(TimeConvert("12:01:00AM"))
}

func TimeConvert(clock string) string {
	res, err := time.Parse("03:04:05PM", clock)
	if err != nil {
		panic(err)
	}

	return res.Format("15:04:05")
}
