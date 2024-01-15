//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLDynamicLibrary
*/

type MTLDynamicLibrary struct {

}

func (r *MTLDynamicLibrary) SerializeToURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDynamicLibrary) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDynamicLibrary) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLDynamicLibrary) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDynamicLibrary) InstallName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

