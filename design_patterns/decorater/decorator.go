package main

type DecoratorInterface interface {
	action() string
}

type Decorator struct {
	dt []DecoratorInterface
}

func (d Decorator) action() string {
	var allStr string
	for _, v := range d.dt {
		allStr += v.action()
	}
	return allStr
}
