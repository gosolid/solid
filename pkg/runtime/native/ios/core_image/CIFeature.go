//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface CoreImage.CIFeature : objc.NSObject
*/

type CIFeature struct {
  *objc.NSObject

}

func (r *CIFeature) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFeature) Bounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

