package main

import (
	"fmt"

	"github.com/hiabhi-cpu/githublogs/internal/user"
)

func main() {
	user := user.User{}
	cfg := config{
		user: user,
	}
	repl(&cfg)
	fmt.Println("Hello")
}
