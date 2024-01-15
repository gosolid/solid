//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSNotificationQueue : objc.NSObject
*/

type NSNotificationQueue struct {
  *objc.NSObject

}

func (r *NSNotificationQueue) InitWithNotificationCenter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotificationQueue) EnqueueNotification() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotificationQueue) DequeueNotificationsMatching() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotificationQueue) DefaultQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

