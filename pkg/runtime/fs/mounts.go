//js:package fs

package fs

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"fmt"
	"path"
	"sort"
	"strings"
)

type locator struct {
	fs   FS
	path string
}

type mount struct {
	*locator
	mountPath  string
	mountPoint string
}

type mounts map[string]*locator

func (m mounts) Resolve(ctx context.Context, p string) (*mount, string, error) {
	candidates := []string{}

	for mountPoint := range m {
		if strings.HasPrefix(p, mountPoint) {
			candidates = append(candidates, mountPoint)
		}
	}

	sort.Slice(candidates, func(i, j int) bool {
		return len(candidates[i]) > len(candidates[j])
	})

	if len(candidates) == 0 {
		return nil, "", NewFSError(ctx, ENOENT, fmt.Sprintf("a mount for path %s does not exist on this filesystem", p))
	}

	mountPoint := candidates[0]
	mnt := m[mountPoint]
	resolvedPath := path.Join(mnt.path, strings.TrimPrefix(p, candidates[0]))

	return &mount{
		locator: &locator{
			fs:   mnt.fs,
			path: resolvedPath,
		},
		mountPath:  mnt.path,
		mountPoint: mountPoint,
	}, mountPoint, nil

}

func (m mount) Reverse(p string) string {
	return path.Join(m.mountPoint, strings.TrimPrefix(p, m.mountPath))
}
