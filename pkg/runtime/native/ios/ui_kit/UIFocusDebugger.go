//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIFocusDebugger : objc.NSObject
*/

type UIFocusDebugger struct {
  *objc.NSObject

}

func (r *UIFocusDebugger) PreferredFocusEnvironmentsForEnvironment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusDebugger) Help() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusDebugger) Status() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusDebugger) CheckFocusabilityForItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusDebugger) SimulateFocusUpdateRequestFromEnvironment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusDebugger) FocusGroupsForEnvironment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

