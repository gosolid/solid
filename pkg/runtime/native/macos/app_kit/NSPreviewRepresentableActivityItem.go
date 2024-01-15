//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSPreviewRepresentableActivityItem
*/

type NSPreviewRepresentableActivityItem struct {

}

func (r *NSPreviewRepresentableActivityItem) ImageProvider() (*foundation.NSItemProvider, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPreviewRepresentableActivityItem) IconProvider() (*foundation.NSItemProvider, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPreviewRepresentableActivityItem) Item() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPreviewRepresentableActivityItem) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

