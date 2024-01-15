//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIDocumentBrowserTransitionController : objc.NSObject
*/

type UIDocumentBrowserTransitionController struct {
  *objc.NSObject

}

func (r *UIDocumentBrowserTransitionController) SetTargetView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentBrowserTransitionController) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentBrowserTransitionController) LoadingProgress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentBrowserTransitionController) SetLoadingProgress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentBrowserTransitionController) TargetView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

