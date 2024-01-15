//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSPressureConfiguration : objc.NSObject
*/

type NSPressureConfiguration struct {
  *objc.NSObject

}

func (r *NSPressureConfiguration) InitWithPressureBehavior() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPressureConfiguration) Set() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPressureConfiguration) PressureBehavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

