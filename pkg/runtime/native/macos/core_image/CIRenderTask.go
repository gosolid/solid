//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CoreImage.CIRenderTask : objc.NSObject
*/

type CIRenderTask struct {
  *objc.NSObject

}

func (r *CIRenderTask) WaitUntilCompletedAndReturnError() (*CIRenderInfo, error) {
  return nil, fmt.Errorf("unimplemented")
}

