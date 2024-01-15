//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreLocation.CLLocationPushServiceExtension
*/

type CLLocationPushServiceExtension struct {

}

func (r *CLLocationPushServiceExtension) ServiceExtensionWillTerminate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationPushServiceExtension) DidReceiveLocationPushPayload() error {
  return fmt.Errorf("unimplemented")
}

