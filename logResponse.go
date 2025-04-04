package main

type LogResponse []struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Description string `json:"description"`
	} `json:"payload"`
}
