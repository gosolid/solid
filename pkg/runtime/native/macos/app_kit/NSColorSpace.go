//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
)

/*
interface AppKit.NSColorSpace : objc.NSObject
*/

type NSColorSpace struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *NSColorSpace) InitWithICCProfileData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) InitWithColorSyncProfile() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) AvailableColorSpacesWithModel() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) CGColorSpace() (*core_graphics.CGColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) ExtendedSRGBColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) DisplayP3ColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) ICCProfileData() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) ColorSpaceModel() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) LocalizedName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) SRGBColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) AdobeRGB1998ColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) GenericGrayColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) DeviceGrayColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) InitWithCGColorSpace() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) ColorSyncProfile() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) NumberOfColorComponents() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) ExtendedGenericGamma22GrayColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) GenericCMYKColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) DeviceRGBColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) DeviceCMYKColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) GenericGamma22GrayColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorSpace) GenericRGBColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

