//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CoreLocation.CLMonitorConfiguration : objc.NSObject
*/

type CLMonitorConfiguration struct {
  *objc.NSObject

}

func (r *CLMonitorConfiguration) ConfigWithMonitorName() (*CLMonitorConfiguration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLMonitorConfiguration) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLMonitorConfiguration) Queue() (**objc.NSObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLMonitorConfiguration) EventHandler() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

