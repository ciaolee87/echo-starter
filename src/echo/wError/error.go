package wError

import (
	"github.com/ciaolee87/echo-starter/src/echo/wJSON"
)

type Error struct {
	error
	wJSON.JSON
	Cb *func()
}
