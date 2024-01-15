//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIDocumentPickerExtensionViewController : UIKit.UIViewController
*/

type UIDocumentPickerExtensionViewController struct {
  *UIViewController

}

func (r *UIDocumentPickerExtensionViewController) DocumentPickerMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerExtensionViewController) OriginalURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerExtensionViewController) ValidTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerExtensionViewController) ProviderIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerExtensionViewController) DocumentStorageURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerExtensionViewController) DismissGrantingAccessToURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerExtensionViewController) PrepareForPresentationInMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

