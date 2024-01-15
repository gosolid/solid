//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIDocumentInteractionController : objc.NSObject
*/

type UIDocumentInteractionController struct {
  *objc.NSObject

}

func (r *UIDocumentInteractionController) PresentOpenInMenuFromRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) DismissPreviewAnimated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) DismissMenuAnimated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) SetAnnotation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) PresentOptionsMenuFromBarButtonItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) SetURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) SetUTI() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) Annotation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) PresentPreviewAnimated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) GestureRecognizers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) PresentOpenInMenuFromBarButtonItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) PresentOptionsMenuFromRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) URL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) UTI() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) Icons() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDocumentInteractionController) InteractionControllerWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

