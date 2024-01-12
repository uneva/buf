package main

import _ "embed"

//go:embed template.gin.tpl
var tpl string

type service struct{}

type method struct{}

func (s *service) execute() string {
	return tpl
}
