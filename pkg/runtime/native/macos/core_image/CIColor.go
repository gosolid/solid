//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreImage.CIColor : objc.NSObject
*/

type CIColor struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CIColor) YellowColor() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) NumberOfComponents() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIColor) Components() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIColor) Alpha() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIColor) GreenColor() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) BlueColor() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) CyanColor() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) RedColor() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) ColorWithCGColor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) ColorWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) InitWithRed() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) Red() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIColor) StringRepresentation() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) GrayColor() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) ColorSpace() (*core_graphics.CGColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) Green() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIColor) BlackColor() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) MagentaColor() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) ClearColor() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) ColorWithRed() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) InitWithCGColor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColor) Blue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIColor) WhiteColor() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

