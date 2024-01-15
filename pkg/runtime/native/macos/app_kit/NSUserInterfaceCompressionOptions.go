//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSUserInterfaceCompressionOptions : objc.NSObject
*/

type NSUserInterfaceCompressionOptions struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSCoding
}

func (r *NSUserInterfaceCompressionOptions) OptionsByAddingOptions() (*NSUserInterfaceCompressionOptions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceCompressionOptions) BreakEqualWidthsOption() (*NSUserInterfaceCompressionOptions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceCompressionOptions) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceCompressionOptions) InitWithCompressionOptions() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceCompressionOptions) ContainsOptions() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceCompressionOptions) IntersectsOptions() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceCompressionOptions) HideTextOption() (*NSUserInterfaceCompressionOptions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceCompressionOptions) ReduceMetricsOption() (*NSUserInterfaceCompressionOptions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceCompressionOptions) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceCompressionOptions) IsEmpty() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceCompressionOptions) HideImagesOption() (*NSUserInterfaceCompressionOptions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceCompressionOptions) StandardOptions() (*NSUserInterfaceCompressionOptions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceCompressionOptions) InitWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceCompressionOptions) OptionsByRemovingOptions() (*NSUserInterfaceCompressionOptions, error) {
  return nil, fmt.Errorf("unimplemented")
}

