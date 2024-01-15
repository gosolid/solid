//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_image"
)

/*
interface AppKit.NSCIImageRep : AppKit.NSImageRep
*/

type NSCIImageRep struct {
  *NSImageRep

}

func (r *NSCIImageRep) ImageRepWithCIImage() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCIImageRep) InitWithCIImage() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCIImageRep) CIImage() (*core_image.CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

