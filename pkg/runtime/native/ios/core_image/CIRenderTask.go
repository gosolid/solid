//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface CoreImage.CIRenderTask : objc.NSObject
*/

type CIRenderTask struct {
  *objc.NSObject

}

func (r *CIRenderTask) WaitUntilCompletedAndReturnError() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

