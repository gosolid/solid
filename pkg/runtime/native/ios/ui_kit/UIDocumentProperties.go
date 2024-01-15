//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIDocumentProperties : objc.NSObject
*/

type UIDocumentProperties struct {
  *objc.NSObject

}

func (r *UIDocumentProperties) InitWithMetadata() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentProperties) Metadata() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentProperties) DragItemsProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentProperties) SetDragItemsProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentProperties) SetWantsIconRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentProperties) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentProperties) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentProperties) InitWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentProperties) SetMetadata() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentProperties) ActivityViewControllerProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentProperties) SetActivityViewControllerProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentProperties) WantsIconRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

