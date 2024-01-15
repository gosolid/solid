//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreLocation.CLPlacemark : objc.NSObject
*/

type CLPlacemark struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *CLPlacemark) AddressDictionary() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) ISOcountryCode() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) InlandWater() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) AreasOfInterest() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) Region() (*CLRegion, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) Country() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) SubThoroughfare() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) AdministrativeArea() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) Locality() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) SubAdministrativeArea() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) PostalCode() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) TimeZone() (*foundation.NSTimeZone, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) Thoroughfare() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) SubLocality() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) Ocean() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) InitWithPlacemark() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLPlacemark) Location() (*CLLocation, error) {
  return nil, fmt.Errorf("unimplemented")
}

