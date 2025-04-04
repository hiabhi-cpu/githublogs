package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hiabhi-cpu/githublogs/internal/user"
)

type config struct {
	user user.User
}

func repl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("github-activity > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) != 1 {
			fmt.Println("Wrong Input")
			continue
		}
		cfg.user.Username = words[0]
		// fmt.Println(words[0])

	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	temp := strings.Split(text, " ")
	res := []string{}
	for _, str := range temp {
		if len(str) != 0 {
			res = append(res, str)
		}
	}
	return res
}
