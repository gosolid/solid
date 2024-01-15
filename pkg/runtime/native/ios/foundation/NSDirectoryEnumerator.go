//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSDirectoryEnumerator : Foundation.NSEnumerator
*/

type NSDirectoryEnumerator struct {
  *NSEnumerator

}

func (r *NSDirectoryEnumerator) DirectoryAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDirectoryEnumerator) IsEnumeratingDirectoryPostOrder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDirectoryEnumerator) Level() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDirectoryEnumerator) SkipDescendents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDirectoryEnumerator) SkipDescendants() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDirectoryEnumerator) FileAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

