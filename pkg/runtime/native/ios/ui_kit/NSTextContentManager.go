//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSTextContentManager : objc.NSObject
*/

type NSTextContentManager struct {
  *objc.NSObject

}

func (r *NSTextContentManager) AddTextLayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) TextElementsForRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) AutomaticallySynchronizesToBackingStore() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) RecordEditActionInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) AutomaticallySynchronizesTextLayoutManagers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) SetAutomaticallySynchronizesTextLayoutManagers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) PrimaryTextLayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) SetAutomaticallySynchronizesToBackingStore() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) TextLayoutManagers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) SetPrimaryTextLayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) RemoveTextLayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) SynchronizeTextLayoutManagers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) PerformEditingTransactionUsingBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentManager) HasEditingTransaction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

