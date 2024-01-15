//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSUserNotificationAction : objc.NSObject
*/

type NSUserNotificationAction struct {
  *objc.NSObject
  *NSCopying
}

func (r *NSUserNotificationAction) ActionWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotificationAction) Identifier() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotificationAction) Title() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

