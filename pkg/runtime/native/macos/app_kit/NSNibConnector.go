//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSNibConnector : objc.NSObject
*/

type NSNibConnector struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSNibConnector) Label() (*objc.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNibConnector) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNibConnector) ReplaceObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNibConnector) EstablishConnection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNibConnector) Source() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNibConnector) SetSource() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNibConnector) Destination() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNibConnector) SetDestination() error {
  return fmt.Errorf("unimplemented")
}

