//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSStoryboardSegue : objc.NSObject
*/

type NSStoryboardSegue struct {
  *objc.NSObject

}

func (r *NSStoryboardSegue) SegueWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStoryboardSegue) InitWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStoryboardSegue) Perform() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStoryboardSegue) Identifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStoryboardSegue) SourceController() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStoryboardSegue) DestinationController() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

