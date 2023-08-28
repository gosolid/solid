package common

import (
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/gosolid/solid/pkg/runtime/fs"
)

func CreateFS(script string) (fs.FS, string, error) {
	if appPath, err := findWorkspaceRoot(script); err != nil {
		return nil, "", err
	} else {
		root := fs.NewFS()
		if err := root.Mount("/app", fs.NewLocalFS(appPath), "/"); err != nil {
			return nil, "", err
		} else {
			return root, "/app" + strings.TrimPrefix(script, appPath), nil
		}
	}
}

func findWorkspaceRoot(script string) (string, error) {
	if stats, err := os.Stat(script); err != nil {
		return "", err
	} else if !stats.IsDir() {
		script = path.Dir(script)
	}

	var workspace *string

	for p := script; ; p = path.Dir(p) {
		if stats, err := os.Stat(path.Join(p, "package.json")); err == nil && !stats.IsDir() {
			if workspace == nil {
				p2 := p
				workspace = &p2
			}

			if b, err := os.ReadFile(path.Join(p, "package.json")); err == nil {
				var descriptionFileData map[string]any
				if err := json.Unmarshal(b, &descriptionFileData); err == nil {
					if workspaces, ok := descriptionFileData["workspaces"]; ok {
						if _, ok := workspaces.(map[string]any); ok {
							return p, nil
						}

						if _, ok := workspaces.([]any); ok {
							return p, nil
						}
					}
				}
			}
		}

		if p == "/" {
			break
		}
	}

	if workspace != nil {
		return *workspace, nil
	} else {
		return script, nil
	}
}
