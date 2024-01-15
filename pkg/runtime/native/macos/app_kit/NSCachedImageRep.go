//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSCachedImageRep : AppKit.NSImageRep
*/

type NSCachedImageRep struct {
  *NSImageRep

}

func (r *NSCachedImageRep) InitWithWindow() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCachedImageRep) InitWithSize() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCachedImageRep) Window() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCachedImageRep) Rect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

