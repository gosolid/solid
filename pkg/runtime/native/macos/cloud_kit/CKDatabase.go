//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CloudKit.CKDatabase : objc.NSObject
*/

type CKDatabase struct {
  *objc.NSObject

}

func (r *CKDatabase) AddOperation() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKDatabase) DatabaseScope() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKDatabase) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKDatabase) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

