package common

import "github.com/usvc/config"

var (
	Config = config.Map{
		"git":   &config.Bool{Shorthand: "g"},
		"apply": &config.Bool{Shorthand: "A"},
	}
)
