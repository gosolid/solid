//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIWindowScene : UIKit.UIScene
*/

type UIWindowScene struct {
  *UIScene

}

func (r *UIWindowScene) RequestGeometryUpdateWithPreferences() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowScene) InterfaceOrientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowScene) Windows() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowScene) ActivityItemsConfigurationSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowScene) Screen() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowScene) CoordinateSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowScene) TraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowScene) EffectiveGeometry() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowScene) SizeRestrictions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowScene) WindowingBehaviors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowScene) KeyWindow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowScene) SetActivityItemsConfigurationSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowScene) IsFullScreen() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

