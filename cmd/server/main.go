package main

import "github.com/mdev5000/vectest2/webb/server"

func main() {
	e, err := server.Serve()
	e.Logger.Fatal(err)
}
