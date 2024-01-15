//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSStringDrawingContext : objc.NSObject
*/

type NSStringDrawingContext struct {
  *objc.NSObject

}

func (r *NSStringDrawingContext) MinimumScaleFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStringDrawingContext) SetMinimumScaleFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStringDrawingContext) ActualScaleFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStringDrawingContext) TotalBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

