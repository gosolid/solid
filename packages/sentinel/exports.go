package sentinel

import (
	"embed"
)

//go:embed all:lib/**
var FS embed.FS
