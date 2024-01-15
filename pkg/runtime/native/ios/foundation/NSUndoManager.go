//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSUndoManager : objc.NSObject
*/

type NSUndoManager struct {
  *objc.NSObject

}

func (r *NSUndoManager) RegisterUndoWithTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) CanRedo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) IsUndoing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) BeginUndoGrouping() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) EndUndoGrouping() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) IsUndoRegistrationEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) SetRunLoopModes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) SetGroupsByEvent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) IsRedoing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) EnableUndoRegistration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) UndoMenuTitleForUndoActionName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) SetLevelsOfUndo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) CanUndo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) UndoActionIsDiscardable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) Undo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) UndoNestedGroup() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) PrepareWithInvocationTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) GroupingLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) UndoMenuItemTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) UndoActionName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) DisableUndoRegistration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) Redo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) SetActionIsDiscardable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) LevelsOfUndo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) RunLoopModes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) RemoveAllActions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) RemoveAllActionsWithTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) SetActionName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) GroupsByEvent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) RedoMenuItemTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) RedoMenuTitleForUndoActionName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) RedoActionIsDiscardable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUndoManager) RedoActionName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

