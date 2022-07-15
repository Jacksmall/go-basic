package main

import "fmt"

type RequestHelper struct {
	desc string
}

type ProcessRequest interface {
	process(RequestHelper) string
}

type DecoratorProcess struct {
	dp []ProcessRequest
}

func (dp *DecoratorProcess) Use(p ProcessRequest) {
	dp.dp = append(dp.dp, p)
}

func (dp *DecoratorProcess) Print(rh RequestHelper) string {
	var sumStr string

	for _, p := range dp.dp {
		sumStr += p.process(rh)
	}

	return sumStr
}

type MainProcess struct{}

func (mp MainProcess) process(RequestHelper) string {
	return "Main"
}

type LogProcess struct{}

func (lp LogProcess) process(req RequestHelper) string {
	return fmt.Sprintln("Logger")
}

type AuthorizationProcess struct{}

func (ap AuthorizationProcess) process(req RequestHelper) string {
	return fmt.Sprintln("Authorization")
}

type StructureProcess struct{}

func (sr StructureProcess) process(req RequestHelper) string {
	return fmt.Sprintln("Structure")
}

func main() {
	req := RequestHelper{
		desc: "index",
	}

	dd := DecoratorProcess{
		dp: make([]ProcessRequest, 0),
	}

	dd.Use(LogProcess{})
	dd.Use(AuthorizationProcess{})
	dd.Use(StructureProcess{})
	dd.Use(MainProcess{})

	fmt.Println(dd.Print(req))
}
