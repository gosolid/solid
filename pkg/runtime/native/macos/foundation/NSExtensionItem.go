//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSExtensionItem : objc.NSObject
*/

type NSExtensionItem struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSExtensionItem) SetAttachments() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSExtensionItem) UserInfo() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExtensionItem) SetUserInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSExtensionItem) AttributedTitle() (*NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExtensionItem) SetAttributedTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSExtensionItem) AttributedContentText() (*NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExtensionItem) SetAttributedContentText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSExtensionItem) Attachments() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

