//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIFont : objc.NSObject
*/

type UIFont struct {
  *objc.NSObject

}

func (r *UIFont) FontNamesForFamilyName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) BoldSystemFontOfSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) FontWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) Leading() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) PreferredFontForTextStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) FamilyNames() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) Descender() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) CapHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) XHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) SystemFontOfSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) MonospacedSystemFontOfSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) FontWithSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) FamilyName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) PointSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) LineHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) FontWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) ItalicSystemFontOfSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) MonospacedDigitSystemFontOfSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) FontName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) Ascender() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFont) FontDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

