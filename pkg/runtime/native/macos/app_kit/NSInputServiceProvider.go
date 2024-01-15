//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSInputServiceProvider
*/

type NSInputServiceProvider struct {

}

func (r *NSInputServiceProvider) MarkedTextSelectionChanged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInputServiceProvider) WantsToInterpretAllKeystrokes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSInputServiceProvider) MarkedTextAbandoned() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInputServiceProvider) Terminate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInputServiceProvider) WantsToHandleMouseEvents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSInputServiceProvider) InputClientResignActive() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInputServiceProvider) ActiveConversationChanged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInputServiceProvider) InsertText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInputServiceProvider) DoCommandBySelector() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInputServiceProvider) CanBeDisabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSInputServiceProvider) WantsToDelayTextChangeNotifications() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSInputServiceProvider) ActiveConversationWillChange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInputServiceProvider) InputClientBecomeActive() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInputServiceProvider) InputClientEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInputServiceProvider) InputClientDisabled() error {
  return fmt.Errorf("unimplemented")
}

