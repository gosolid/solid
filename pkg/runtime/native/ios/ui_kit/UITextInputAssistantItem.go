//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITextInputAssistantItem : objc.NSObject
*/

type UITextInputAssistantItem struct {
  *objc.NSObject

}

func (r *UITextInputAssistantItem) SetAllowsHidingShortcuts() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputAssistantItem) LeadingBarButtonGroups() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputAssistantItem) SetLeadingBarButtonGroups() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputAssistantItem) TrailingBarButtonGroups() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputAssistantItem) SetTrailingBarButtonGroups() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputAssistantItem) AllowsHidingShortcuts() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

