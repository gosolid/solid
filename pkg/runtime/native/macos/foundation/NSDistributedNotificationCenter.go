//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSDistributedNotificationCenter : Foundation.NSNotificationCenter
*/

type NSDistributedNotificationCenter struct {
  *NSNotificationCenter

}

func (r *NSDistributedNotificationCenter) Suspended() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDistributedNotificationCenter) SetSuspended() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDistributedNotificationCenter) NotificationCenterForType() (*NSDistributedNotificationCenter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDistributedNotificationCenter) DefaultCenter() (*NSDistributedNotificationCenter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDistributedNotificationCenter) AddObserver() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDistributedNotificationCenter) PostNotificationName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDistributedNotificationCenter) RemoveObserver() error {
  return fmt.Errorf("unimplemented")
}

