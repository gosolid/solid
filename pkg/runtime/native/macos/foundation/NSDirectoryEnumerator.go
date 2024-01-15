//js:package native/macos/foundation
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

func (r *NSDirectoryEnumerator) DirectoryAttributes() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDirectoryEnumerator) IsEnumeratingDirectoryPostOrder() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDirectoryEnumerator) Level() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDirectoryEnumerator) SkipDescendents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDirectoryEnumerator) SkipDescendants() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDirectoryEnumerator) FileAttributes() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

