package main

import "github.com/mdev5000/vectest2/webf"

func main() {
	if err := webf.Main(); err != nil {
		panic(err)
	}
}
