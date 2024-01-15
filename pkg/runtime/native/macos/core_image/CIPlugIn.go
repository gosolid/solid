//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CoreImage.CIPlugIn : objc.NSObject
*/

type CIPlugIn struct {
  *objc.NSObject

}

func (r *CIPlugIn) LoadNonExecutablePlugIn() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPlugIn) LoadAllPlugIns() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPlugIn) LoadNonExecutablePlugIns() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPlugIn) LoadPlugIn() error {
  return fmt.Errorf("unimplemented")
}

