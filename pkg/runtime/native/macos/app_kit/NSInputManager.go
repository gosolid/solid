//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSInputManager : objc.NSObject
*/

type NSInputManager struct {
  *objc.NSObject
  *NSTextInput
}

func (r *NSInputManager) CycleToNextInputServerInLanguage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInputManager) LocalizedInputManagerName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInputManager) MarkedTextSelectionChanged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInputManager) WantsToInterpretAllKeystrokes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSInputManager) Language() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInputManager) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInputManager) CycleToNextInputLanguage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInputManager) InitWithName() (*NSInputManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInputManager) MarkedTextAbandoned() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInputManager) HandleMouseEvent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSInputManager) CurrentInputManager() (*NSInputManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInputManager) Server() (*NSInputServer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInputManager) WantsToHandleMouseEvents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSInputManager) WantsToDelayTextChangeNotifications() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

