//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIScene : UIKit.UIResponder
*/

type UIScene struct {
  *UIResponder

}

func (r *UIScene) InitWithSession() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScene) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScene) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScene) OpenURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScene) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScene) SetSubtitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScene) ActivationState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScene) Subtitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScene) ActivationConditions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScene) SetActivationConditions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScene) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScene) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScene) Session() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScene) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

