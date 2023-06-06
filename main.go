package main

import "strings"

type Logger interface {
	Log(msg string) error
}
type user struct {
	isAngry bool
	logger  Logger
}

func (u user) Talk(sentence string) error {
	if u.isAngry {
		return u.logger.Log(strings.ToUpper(sentence))
	}
	return u.logger.Log(sentence)
}
func main() {}
