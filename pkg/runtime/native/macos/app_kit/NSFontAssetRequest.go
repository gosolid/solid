//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSFontAssetRequest : objc.NSObject
*/

type NSFontAssetRequest struct {
  *objc.NSObject
  *foundation.NSProgressReporting
}

func (r *NSFontAssetRequest) DownloadedFontDescriptors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontAssetRequest) Progress() (*foundation.NSProgress, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontAssetRequest) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontAssetRequest) InitWithFontDescriptors() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontAssetRequest) DownloadFontAssetsWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

