//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKAsset : objc.NSObject
*/

type CKAsset struct {
  *objc.NSObject

}

func (r *CKAsset) InitWithFileURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKAsset) FileURL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKAsset) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKAsset) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

