//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSDataAsset : objc.NSObject
*/

type NSDataAsset struct {
  *objc.NSObject

}

func (r *NSDataAsset) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDataAsset) InitWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDataAsset) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDataAsset) Data() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDataAsset) TypeIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

