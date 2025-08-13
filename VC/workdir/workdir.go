package workdir

import "strings"

// you can use this library freely: "github.com/otiai10/copy"

type WorkDir struct {
	root string
	files map[string]string
}
