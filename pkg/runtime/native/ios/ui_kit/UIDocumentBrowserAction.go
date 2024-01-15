//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIDocumentBrowserAction : objc.NSObject
*/

type UIDocumentBrowserAction struct {
  *objc.NSObject

}

func (r *UIDocumentBrowserAction) LocalizedTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentBrowserAction) Availability() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentBrowserAction) Image() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentBrowserAction) SupportedContentTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentBrowserAction) SetSupportsMultipleItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentBrowserAction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentBrowserAction) InitWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentBrowserAction) SetSupportedContentTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentBrowserAction) SupportsMultipleItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentBrowserAction) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentBrowserAction) SetImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

