package main

type Authorization struct{}

func (a Authorization) action() string {
	return "authorize"
}
