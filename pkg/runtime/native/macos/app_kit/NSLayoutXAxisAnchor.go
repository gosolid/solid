//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSLayoutXAxisAnchor : AppKit.NSLayoutAnchor
*/

type NSLayoutXAxisAnchor struct {
  *NSLayoutAnchor

}

func (r *NSLayoutXAxisAnchor) AnchorWithOffsetToAnchor() (*NSLayoutDimension, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutXAxisAnchor) ConstraintEqualToSystemSpacingAfterAnchor() (*NSLayoutConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutXAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor() (*NSLayoutConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutXAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingAfterAnchor() (*NSLayoutConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

