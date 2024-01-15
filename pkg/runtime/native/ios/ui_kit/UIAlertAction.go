//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIAlertAction : objc.NSObject
*/

type UIAlertAction struct {
  *objc.NSObject

}

func (r *UIAlertAction) ActionWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertAction) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertAction) Style() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertAction) IsEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertAction) SetEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

