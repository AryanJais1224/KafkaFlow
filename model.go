package main

type Event struct {
	ID      string `json:"id"`
	Data    string `json:"data"`
	Retry   int    `json:"retry"`
}
