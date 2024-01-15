//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CoreImage.CIDetector : objc.NSObject
*/

type CIDetector struct {
  *objc.NSObject

}

func (r *CIDetector) DetectorOfType() (*CIDetector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDetector) FeaturesInImage() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

