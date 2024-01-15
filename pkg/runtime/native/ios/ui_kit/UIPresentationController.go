//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPresentationController : objc.NSObject
*/

type UIPresentationController struct {
  *objc.NSObject

}

func (r *UIPresentationController) ContainerViewWillLayoutSubviews() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) DismissalTransitionDidEnd() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) FrameOfPresentedViewInContainerView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) ShouldRemovePresentersView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) SetOverrideTraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) ContainerViewDidLayoutSubviews() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) PresentationTransitionDidEnd() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) ContainerView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) OverrideTraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) PresentedViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) PresentationStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) AdaptivePresentationStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) PresentedView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) ShouldPresentInFullscreen() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) InitWithPresentedViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) AdaptivePresentationStyleForTraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) PresentationTransitionWillBegin() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) DismissalTransitionWillBegin() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPresentationController) PresentingViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

