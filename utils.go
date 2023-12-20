package main

import (
	"bytes"
	"context"

	"github.com/a-h/templ"
)

func Render(component templ.Component) []byte {
	buf := new(bytes.Buffer)
	component.Render(context.Background(), buf)
	return buf.Bytes()
}
