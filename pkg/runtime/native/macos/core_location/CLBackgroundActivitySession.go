//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CoreLocation.CLBackgroundActivitySession : objc.NSObject
*/

type CLBackgroundActivitySession struct {
  *objc.NSObject

}

func (r *CLBackgroundActivitySession) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBackgroundActivitySession) Invalidate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLBackgroundActivitySession) BackgroundActivitySession() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBackgroundActivitySession) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

