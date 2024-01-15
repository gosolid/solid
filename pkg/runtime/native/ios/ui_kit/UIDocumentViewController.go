//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIDocumentViewController : UIKit.UIViewController
*/

type UIDocumentViewController struct {
  *UIViewController

}

func (r *UIDocumentViewController) UndoRedoItemGroup() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentViewController) InitWithDocument() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentViewController) NavigationItemDidUpdate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentViewController) OpenDocumentWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentViewController) DocumentDidOpen() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentViewController) Document() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentViewController) SetDocument() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

