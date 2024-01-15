//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSPageControllerDelegate
*/

type NSPageControllerDelegate struct {

}

func (r *NSPageControllerDelegate) PageController() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPageControllerDelegate) PageControllerWillStartLiveTransition() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPageControllerDelegate) PageControllerDidEndLiveTransition() error {
  return fmt.Errorf("unimplemented")
}

