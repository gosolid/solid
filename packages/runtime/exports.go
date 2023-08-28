package runtime

import (
	"embed"
)

//go:embed all:lib/**
var FS embed.FS
