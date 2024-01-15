//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSMovie : objc.NSObject
*/

type NSMovie struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSMovie) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMovie) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMovie) InitWithMovie() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMovie) QTMovie() (*QTMovie, error) {
  return nil, fmt.Errorf("unimplemented")
}

