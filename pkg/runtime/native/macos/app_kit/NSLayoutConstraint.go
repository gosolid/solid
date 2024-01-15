//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSLayoutConstraint : objc.NSObject
*/

type NSLayoutConstraint struct {
  *objc.NSObject

}

func (r *NSLayoutConstraint) SecondAnchor() (*NSLayoutAnchor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) Multiplier() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) Constant() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) ConstraintsWithVisualFormat() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) ShouldBeArchived() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) FirstItem() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) FirstAttribute() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) SetActive() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) ConstraintWithItem() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) FirstAnchor() (*NSLayoutAnchor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) Relation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) IsActive() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) ActivateConstraints() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) Priority() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) SecondItem() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) SetConstant() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) DeactivateConstraints() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) SetPriority() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) SetShouldBeArchived() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) SecondAttribute() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

