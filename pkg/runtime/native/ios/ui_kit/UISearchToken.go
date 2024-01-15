//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISearchToken : objc.NSObject
*/

type UISearchToken struct {
  *objc.NSObject

}

func (r *UISearchToken) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchToken) TokenWithIcon() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchToken) RepresentedObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchToken) SetRepresentedObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchToken) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

