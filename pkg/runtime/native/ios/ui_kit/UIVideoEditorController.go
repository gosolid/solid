//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIVideoEditorController : UIKit.UINavigationController
*/

type UIVideoEditorController struct {
  *UINavigationController

}

func (r *UIVideoEditorController) SetVideoMaximumDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIVideoEditorController) VideoQuality() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIVideoEditorController) SetVideoQuality() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIVideoEditorController) VideoMaximumDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIVideoEditorController) CanEditVideoAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIVideoEditorController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIVideoEditorController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIVideoEditorController) VideoPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIVideoEditorController) SetVideoPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

