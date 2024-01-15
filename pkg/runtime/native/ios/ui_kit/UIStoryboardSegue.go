//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIStoryboardSegue : objc.NSObject
*/

type UIStoryboardSegue struct {
  *objc.NSObject

}

func (r *UIStoryboardSegue) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStoryboardSegue) SourceViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStoryboardSegue) DestinationViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStoryboardSegue) SegueWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStoryboardSegue) InitWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStoryboardSegue) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStoryboardSegue) Perform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

