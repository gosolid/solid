//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSLayoutDimension : AppKit.NSLayoutAnchor
*/

type NSLayoutDimension struct {
  *NSLayoutAnchor

}

func (r *NSLayoutDimension) ConstraintLessThanOrEqualToConstant() (*NSLayoutConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutDimension) ConstraintEqualToAnchor() (*NSLayoutConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutDimension) ConstraintGreaterThanOrEqualToAnchor() (*NSLayoutConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutDimension) ConstraintLessThanOrEqualToAnchor() (*NSLayoutConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutDimension) ConstraintEqualToConstant() (*NSLayoutConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutDimension) ConstraintGreaterThanOrEqualToConstant() (*NSLayoutConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

