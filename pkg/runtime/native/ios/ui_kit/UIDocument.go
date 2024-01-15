//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIDocument : objc.NSObject
*/

type UIDocument struct {
  *objc.NSObject

}

func (r *UIDocument) FileModificationDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) EnableEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) RevertToContentsOfURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) DocumentState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) Progress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) FinishedHandlingError() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) ReadFromURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) PerformAsynchronousFileAccessUsingBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) HandleError() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) UserInteractionNoLongerPermittedForError() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) LocalizedName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) UndoManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) HasUnsavedChanges() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) FileAttributesToWriteToURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) SavingFileType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) CloseWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) UpdateChangeCountWithToken() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) SaveToURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) WriteContents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) InitWithFileURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) AutosaveWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) DisableEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) FileNameExtensionForType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) FileURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) ContentsForType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) SetFileModificationDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) SetUndoManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) ChangeCountTokenForSaveOperation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) LoadFromContents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) UpdateChangeCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) FileType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocument) OpenWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

