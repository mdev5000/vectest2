package main

import (
	"fmt"
	"github.com/mdev5000/vectest2/proto"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	user := &proto.User{
		Id:    "user1",
		Name:  "User 1",
		Email: "user1@test.com",
	}
	userb, err := user.Marshal()
	if err != nil {
		return err
	}

	user2 := &proto.User{}
	if err := user2.Unmarshal(userb); err != nil {
		return err
	}

	fmt.Println(user2.GetId())
	fmt.Println(user2.GetName())
	fmt.Println(user2.GetEmail())

	return nil
}
