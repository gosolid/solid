//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSUndoManager : objc.NSObject
*/

type NSUndoManager struct {
  *objc.NSObject

}

func (r *NSUndoManager) UndoNestedGroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) RemoveAllActions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) RemoveAllActionsWithTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) SetGroupsByEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) CanUndo() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) UndoActionName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) BeginUndoGrouping() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) Redo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) UndoMenuTitleForUndoActionName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) RedoActionIsDiscardable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) SetActionName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) GroupsByEvent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) RegisterUndoWithTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) SetLevelsOfUndo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) SetRunLoopModes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) RedoMenuItemTitle() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) EndUndoGrouping() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) EnableUndoRegistration() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) PrepareWithInvocationTarget() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) SetActionIsDiscardable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) LevelsOfUndo() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) RedoActionName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) UndoMenuItemTitle() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) GroupingLevel() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) CanRedo() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) IsUndoing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) RunLoopModes() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) IsRedoing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) DisableUndoRegistration() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) Undo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) RedoMenuTitleForUndoActionName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) IsUndoRegistrationEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) UndoActionIsDiscardable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

