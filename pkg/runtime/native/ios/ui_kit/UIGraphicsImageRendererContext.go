//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIGraphicsImageRendererContext : UIKit.UIGraphicsRendererContext
*/

type UIGraphicsImageRendererContext struct {
  *UIGraphicsRendererContext

}

func (r *UIGraphicsImageRendererContext) CurrentImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

