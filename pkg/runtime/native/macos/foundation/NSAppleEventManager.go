//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSAppleEventManager : objc.NSObject
*/

type NSAppleEventManager struct {
  *objc.NSObject

}

func (r *NSAppleEventManager) CurrentAppleEvent() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventManager) CurrentReplyAppleEvent() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventManager) SharedAppleEventManager() (*NSAppleEventManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventManager) RemoveEventHandlerForEventClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAppleEventManager) AppleEventForSuspensionID() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventManager) SetCurrentAppleEventAndReplyEventWithSuspensionID() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAppleEventManager) ResumeWithSuspensionID() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAppleEventManager) SetEventHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAppleEventManager) DispatchRawAppleEvent() (int16, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventManager) SuspendCurrentAppleEvent() (*NSAppleEventManagerSuspension, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventManager) ReplyAppleEventForSuspensionID() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

