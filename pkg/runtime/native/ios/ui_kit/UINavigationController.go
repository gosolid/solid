//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UINavigationController : UIKit.UIViewController
*/

type UINavigationController struct {
  *UIViewController

}

func (r *UINavigationController) VisibleViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) Toolbar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) InitWithNibName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) PopViewControllerAnimated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) HidesBarsWhenKeyboardAppears() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) HidesBarsWhenVerticallyCompact() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) ShowViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) ViewControllers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) SetHidesBarsOnTap() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) SetNavigationBarHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) NavigationBar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) PopToRootViewControllerAnimated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) IsToolbarHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) BarHideOnSwipeGestureRecognizer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) InitWithNavigationBarClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) IsNavigationBarHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) SetToolbarHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) TopViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) SetHidesBarsWhenKeyboardAppears() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) InitWithRootViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) PopToViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) HidesBarsOnSwipe() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) SetHidesBarsWhenVerticallyCompact() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) HidesBarsOnTap() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) PushViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) SetViewControllers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) BarHideOnTapGestureRecognizer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) InteractivePopGestureRecognizer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationController) SetHidesBarsOnSwipe() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

