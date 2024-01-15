//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIPageViewController : UIKit.UIViewController
*/

type UIPageViewController struct {
  *UIViewController

}

func (r *UIPageViewController) DataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageViewController) TransitionStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageViewController) GestureRecognizers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageViewController) IsDoubleSided() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageViewController) SetDoubleSided() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageViewController) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageViewController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageViewController) SpineLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageViewController) InitWithTransitionStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageViewController) SetViewControllers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageViewController) ViewControllers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageViewController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageViewController) SetDataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageViewController) NavigationOrientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

