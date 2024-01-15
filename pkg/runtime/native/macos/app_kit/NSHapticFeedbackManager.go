//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSHapticFeedbackManager : objc.NSObject
*/

type NSHapticFeedbackManager struct {
  *objc.NSObject

}

func (r *NSHapticFeedbackManager) DefaultPerformer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

