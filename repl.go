package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

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
		if words[0] == "exit" {
			break
		}

		cfg.user.Username = words[0]

		url := "https://api.github.com/users/<username>/events"
		url = strings.Replace(url, "<username>", cfg.user.Username, len(cfg.user.Username))
		fmt.Println(url)

		client := http.Client{
			Timeout: time.Minute,
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Errorf("An erroe", err)
			continue
		}

		res, err := client.Do(req)
		if err != nil {
			fmt.Errorf("An erroe", err)
			continue
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			fmt.Println("User not exists")
			continue
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Errorf("An erroe", err)
			continue
		}

		logResponse := LogResponse{}

		if err = json.Unmarshal(data, &logResponse); err != nil {
			fmt.Errorf("Error in data reading", err)
			continue
		}
		for _, r := range logResponse {
			fmt.Println(r.Repo.Name, " : ", r.Type)
		}
		// fmt.Println(logResponse)
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
