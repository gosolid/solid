//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLStitchedLibraryDescriptor : objc.NSObject
*/

type MTLStitchedLibraryDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLStitchedLibraryDescriptor) FunctionGraphs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStitchedLibraryDescriptor) SetFunctionGraphs() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLStitchedLibraryDescriptor) Functions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStitchedLibraryDescriptor) SetFunctions() error {
  return fmt.Errorf("unimplemented")
}

