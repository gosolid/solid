//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIDocumentMenuViewController : UIKit.UIViewController
*/

type UIDocumentMenuViewController struct {
  *UIViewController

}

func (r *UIDocumentMenuViewController) InitWithDocumentTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentMenuViewController) InitWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentMenuViewController) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentMenuViewController) AddOptionWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentMenuViewController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentMenuViewController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

