//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSLayoutConstraint : objc.NSObject
*/

type NSLayoutConstraint struct {
  *objc.NSObject

}

func (r *NSLayoutConstraint) DeactivateConstraints() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) Priority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) SetPriority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) SecondItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) FirstAttribute() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) Relation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) ConstraintWithItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) ShouldBeArchived() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) SetShouldBeArchived() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) SecondAttribute() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) FirstAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) SecondAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) Constant() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) ActivateConstraints() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) IsActive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) FirstItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) Multiplier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) SetConstant() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) SetActive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutConstraint) ConstraintsWithVisualFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

