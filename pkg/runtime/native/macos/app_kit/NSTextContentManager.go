//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTextContentManager : objc.NSObject
*/

type NSTextContentManager struct {
  *objc.NSObject
  *NSTextElementProvider
  *foundation.NSSecureCoding
}

func (r *NSTextContentManager) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) AddTextLayoutManager() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) TextElementsForRange() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) AutomaticallySynchronizesTextLayoutManagers() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) SetAutomaticallySynchronizesToBackingStore() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) SynchronizeTextLayoutManagers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) PerformEditingTransactionUsingBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) PrimaryTextLayoutManager() (*NSTextLayoutManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) SetPrimaryTextLayoutManager() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) HasEditingTransaction() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) SetAutomaticallySynchronizesTextLayoutManagers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) AutomaticallySynchronizesToBackingStore() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) TextLayoutManagers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) RemoveTextLayoutManager() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) RecordEditActionInRange() error {
  return fmt.Errorf("unimplemented")
}

