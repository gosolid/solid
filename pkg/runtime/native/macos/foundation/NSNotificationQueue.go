//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSNotificationQueue : objc.NSObject
*/

type NSNotificationQueue struct {
  *objc.NSObject

}

func (r *NSNotificationQueue) DefaultQueue() (*NSNotificationQueue, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotificationQueue) InitWithNotificationCenter() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotificationQueue) EnqueueNotification() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNotificationQueue) DequeueNotificationsMatching() error {
  return fmt.Errorf("unimplemented")
}

