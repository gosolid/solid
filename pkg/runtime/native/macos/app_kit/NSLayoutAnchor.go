//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSLayoutAnchor : objc.NSObject
*/

type NSLayoutAnchor struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSCoding
}

func (r *NSLayoutAnchor) ConstraintEqualToAnchor() (*NSLayoutConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutAnchor) ConstraintGreaterThanOrEqualToAnchor() (*NSLayoutConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutAnchor) ConstraintLessThanOrEqualToAnchor() (*NSLayoutConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutAnchor) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutAnchor) Item() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutAnchor) HasAmbiguousLayout() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLayoutAnchor) ConstraintsAffectingLayout() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

