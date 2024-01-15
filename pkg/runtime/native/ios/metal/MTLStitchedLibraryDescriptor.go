//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLStitchedLibraryDescriptor : objc.NSObject
*/

type MTLStitchedLibraryDescriptor struct {
  *objc.NSObject

}

func (r *MTLStitchedLibraryDescriptor) FunctionGraphs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStitchedLibraryDescriptor) SetFunctionGraphs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStitchedLibraryDescriptor) Functions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStitchedLibraryDescriptor) SetFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

