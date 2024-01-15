//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UITabBarController : UIKit.UIViewController
*/

type UITabBarController struct {
  *UIViewController

}

func (r *UITabBarController) CustomizableViewControllers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarController) TabBar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarController) SelectedViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarController) SetSelectedViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarController) SetSelectedIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarController) MoreNavigationController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarController) SetCustomizableViewControllers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarController) SetViewControllers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarController) ViewControllers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarController) SelectedIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

