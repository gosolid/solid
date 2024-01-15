//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSFilePromiseProvider : objc.NSObject
*/

type NSFilePromiseProvider struct {
  *objc.NSObject
  *NSPasteboardWriting
}

func (r *NSFilePromiseProvider) FileType() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFilePromiseProvider) SetFileType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePromiseProvider) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFilePromiseProvider) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePromiseProvider) UserInfo() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFilePromiseProvider) SetUserInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePromiseProvider) InitWithFileType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFilePromiseProvider) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

