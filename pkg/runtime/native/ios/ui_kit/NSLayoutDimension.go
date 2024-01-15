//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.NSLayoutDimension : UIKit.NSLayoutAnchor
*/

type NSLayoutDimension struct {
  *NSLayoutAnchor

}

func (r *NSLayoutDimension) ConstraintGreaterThanOrEqualToAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutDimension) ConstraintLessThanOrEqualToAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutDimension) ConstraintEqualToConstant() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutDimension) ConstraintGreaterThanOrEqualToConstant() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutDimension) ConstraintLessThanOrEqualToConstant() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutDimension) ConstraintEqualToAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

