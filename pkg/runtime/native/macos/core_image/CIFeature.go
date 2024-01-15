//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreImage.CIFeature : objc.NSObject
*/

type CIFeature struct {
  *objc.NSObject

}

func (r *CIFeature) Bounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIFeature) Type() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

