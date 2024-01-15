//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIDocumentPickerViewController : UIKit.UIViewController
*/

type UIDocumentPickerViewController struct {
  *UIViewController

}

func (r *UIDocumentPickerViewController) AllowsMultipleSelection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerViewController) DirectoryURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerViewController) InitWithURLs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerViewController) InitForExportingURLs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerViewController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerViewController) InitWithDocumentTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerViewController) SetDirectoryURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerViewController) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerViewController) ShouldShowFileExtensions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerViewController) SetShouldShowFileExtensions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerViewController) DocumentPickerMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerViewController) SetAllowsMultipleSelection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerViewController) InitForOpeningContentTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerViewController) InitWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentPickerViewController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

