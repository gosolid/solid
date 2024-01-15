//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CoreLocation.CLLocationUpdater : objc.NSObject
*/

type CLLocationUpdater struct {
  *objc.NSObject

}

func (r *CLLocationUpdater) Pause() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationUpdater) Invalidate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationUpdater) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLLocationUpdater) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLLocationUpdater) LiveUpdaterWithQueue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLLocationUpdater) LiveUpdaterWithConfiguration() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLLocationUpdater) Resume() error {
  return fmt.Errorf("unimplemented")
}

