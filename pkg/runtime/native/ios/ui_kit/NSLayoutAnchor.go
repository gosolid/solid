//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSLayoutAnchor : objc.NSObject
*/

type NSLayoutAnchor struct {
  *objc.NSObject

}

func (r *NSLayoutAnchor) ConstraintEqualToAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutAnchor) ConstraintGreaterThanOrEqualToAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutAnchor) ConstraintLessThanOrEqualToAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutAnchor) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutAnchor) Item() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutAnchor) HasAmbiguousLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutAnchor) ConstraintsAffectingLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

