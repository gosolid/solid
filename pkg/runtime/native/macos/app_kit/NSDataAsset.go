//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSDataAsset : objc.NSObject
*/

type NSDataAsset struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSDataAsset) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDataAsset) InitWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDataAsset) Name() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDataAsset) Data() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDataAsset) TypeIdentifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

