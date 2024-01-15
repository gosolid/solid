//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSFontDescriptor : objc.NSObject
*/

type NSFontDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *NSFontDescriptor) FontDescriptorWithName() (*NSFontDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) FontDescriptorWithMatrix() (*NSFontDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) FontDescriptorWithFace() (*NSFontDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) FontDescriptorWithDesign() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) FontAttributes() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) FontDescriptorWithFontAttributes() (*NSFontDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) MatchingFontDescriptorWithMandatoryKeys() (*NSFontDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) FontDescriptorWithFamily() (*NSFontDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) PointSize() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) Matrix() (*foundation.NSAffineTransform, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) SymbolicTraits() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) ObjectForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) MatchingFontDescriptorsWithMandatoryKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) FontDescriptorByAddingAttributes() (*NSFontDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) RequiresFontAssetRequest() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) InitWithFontAttributes() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) FontDescriptorWithSymbolicTraits() (*NSFontDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) FontDescriptorWithSize() (*NSFontDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontDescriptor) PostscriptName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

