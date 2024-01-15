//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface CoreImage.CIDetector : objc.NSObject
*/

type CIDetector struct {
  *objc.NSObject

}

func (r *CIDetector) DetectorOfType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDetector) FeaturesInImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

