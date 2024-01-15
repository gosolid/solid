//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSLayoutYAxisAnchor : AppKit.NSLayoutAnchor
*/

type NSLayoutYAxisAnchor struct {
  *NSLayoutAnchor

}

func (r *NSLayoutYAxisAnchor) AnchorWithOffsetToAnchor() (*NSLayoutDimension, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutYAxisAnchor) ConstraintEqualToSystemSpacingBelowAnchor() (*NSLayoutConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutYAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor() (*NSLayoutConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutYAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingBelowAnchor() (*NSLayoutConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

