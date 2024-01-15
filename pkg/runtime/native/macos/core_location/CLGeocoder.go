//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CoreLocation.CLGeocoder : objc.NSObject
*/

type CLGeocoder struct {
  *objc.NSObject

}

func (r *CLGeocoder) ReverseGeocodeLocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLGeocoder) GeocodeAddressDictionary() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLGeocoder) GeocodeAddressString() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLGeocoder) CancelGeocode() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLGeocoder) IsGeocoding() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

