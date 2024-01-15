//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSCustomImageRep : AppKit.NSImageRep
*/

type NSCustomImageRep struct {
  *NSImageRep

}

func (r *NSCustomImageRep) InitWithSize() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCustomImageRep) InitWithDrawSelector() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCustomImageRep) DrawingHandler() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCustomImageRep) DrawSelector() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCustomImageRep) Delegate() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

