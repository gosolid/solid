//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSNotificationCenter : objc.NSObject
*/

type NSNotificationCenter struct {
  *objc.NSObject

}

func (r *NSNotificationCenter) RemoveObserver() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotificationCenter) AddObserverForName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotificationCenter) DefaultCenter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotificationCenter) AddObserver() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotificationCenter) PostNotification() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotificationCenter) PostNotificationName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

