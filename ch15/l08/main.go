package main

type email struct {
	recipient string
	subject   string
}

type receipt struct {
	amount int
}

type inbox struct {
	events []any
}

func (i inbox) Select[T any]() []T {
	result := []T{}
	for _, e := range i.events {
		if v, ok := e.(T); ok {
			result = append(result, v)
		}
	}
	return result
}
