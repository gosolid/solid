//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSNotificationCenter : objc.NSObject
*/

type NSNotificationCenter struct {
  *objc.NSObject

}

func (r *NSNotificationCenter) PostNotification() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNotificationCenter) PostNotificationName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNotificationCenter) RemoveObserver() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNotificationCenter) AddObserverForName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotificationCenter) DefaultCenter() (*NSNotificationCenter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotificationCenter) AddObserver() error {
  return fmt.Errorf("unimplemented")
}

