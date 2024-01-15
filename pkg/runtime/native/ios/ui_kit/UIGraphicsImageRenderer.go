//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIGraphicsImageRenderer : UIKit.UIGraphicsRenderer
*/

type UIGraphicsImageRenderer struct {
  *UIGraphicsRenderer

}

func (r *UIGraphicsImageRenderer) JPEGDataWithCompressionQuality() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsImageRenderer) InitWithSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsImageRenderer) InitWithBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsImageRenderer) ImageWithActions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsImageRenderer) PNGDataWithActions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

