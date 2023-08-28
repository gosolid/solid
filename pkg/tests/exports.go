package tests

import (
	"embed"

	"github.com/gosolid/solid/pkg/runtime/fs"
)

//go:embed all:lib/**
var testsFS embed.FS

var FS = fs.NewFS().Mount("/", fs.NewEmbedFS(&testsFS), "/lib")
