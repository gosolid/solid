//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIDropProposal : objc.NSObject
*/

type UIDropProposal struct {
  *objc.NSObject

}

func (r *UIDropProposal) SetPrecise() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDropProposal) PrefersFullSizePreview() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDropProposal) SetPrefersFullSizePreview() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDropProposal) InitWithDropOperation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDropProposal) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDropProposal) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDropProposal) Operation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDropProposal) IsPrecise() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

