//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISheetPresentationControllerDetent : objc.NSObject
*/

type UISheetPresentationControllerDetent struct {
  *objc.NSObject

}

func (r *UISheetPresentationControllerDetent) LargeDetent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISheetPresentationControllerDetent) CustomDetentWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISheetPresentationControllerDetent) ResolvedValueInContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISheetPresentationControllerDetent) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISheetPresentationControllerDetent) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISheetPresentationControllerDetent) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISheetPresentationControllerDetent) MediumDetent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

