//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSShadow : objc.NSObject
*/

type NSShadow struct {
  *objc.NSObject

}

func (r *NSShadow) ShadowBlurRadius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSShadow) SetShadowBlurRadius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSShadow) ShadowColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSShadow) SetShadowColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSShadow) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSShadow) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSShadow) ShadowOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSShadow) SetShadowOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

