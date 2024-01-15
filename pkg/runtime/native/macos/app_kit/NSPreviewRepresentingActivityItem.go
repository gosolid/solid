//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSPreviewRepresentingActivityItem : objc.NSObject
*/

type NSPreviewRepresentingActivityItem struct {
  *objc.NSObject
  *NSPreviewRepresentableActivityItem
}

func (r *NSPreviewRepresentingActivityItem) InitWithItem() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPreviewRepresentingActivityItem) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPreviewRepresentingActivityItem) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

