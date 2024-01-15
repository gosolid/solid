//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIFontDescriptor : objc.NSObject
*/

type UIFontDescriptor struct {
  *objc.NSObject

}

func (r *UIFontDescriptor) FontDescriptorWithSymbolicTraits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) FontDescriptorWithDesign() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) PointSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) Matrix() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) FontDescriptorWithFontAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) PreferredFontDescriptorWithTextStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) FontDescriptorWithSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) FontDescriptorWithFace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) FontDescriptorWithMatrix() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) ObjectForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) MatchingFontDescriptorsWithMandatoryKeys() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) InitWithFontAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) FontDescriptorWithFamily() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) SymbolicTraits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) FontAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) FontDescriptorWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) FontDescriptorByAddingAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontDescriptor) PostscriptName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

