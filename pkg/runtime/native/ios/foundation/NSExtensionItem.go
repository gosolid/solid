//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSExtensionItem : objc.NSObject
*/

type NSExtensionItem struct {
  *objc.NSObject

}

func (r *NSExtensionItem) SetUserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExtensionItem) AttributedTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExtensionItem) SetAttributedTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExtensionItem) AttributedContentText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExtensionItem) SetAttributedContentText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExtensionItem) Attachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExtensionItem) SetAttachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExtensionItem) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

