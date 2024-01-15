//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/foundation"
  "fmt"
)

/*
interface UIKit.UIActivityItemProvider : Foundation.NSOperation
*/

type UIActivityItemProvider struct {
  *foundation.NSOperation

}

func (r *UIActivityItemProvider) PlaceholderItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemProvider) ActivityType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemProvider) Item() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemProvider) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemProvider) InitWithPlaceholderItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

