//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface CoreImage.CIColor : objc.NSObject
*/

type CIColor struct {
  *objc.NSObject

}

func (r *CIColor) ColorWithCGColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) ColorWithRed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) ColorWithString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) InitWithCGColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) StringRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) BlackColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) GreenColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) YellowColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) Alpha() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) ColorSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) Green() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) GrayColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) CyanColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) ClearColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) WhiteColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) RedColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) BlueColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) InitWithRed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) NumberOfComponents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) Components() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) Red() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) Blue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) MagentaColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

