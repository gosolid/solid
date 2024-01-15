//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CloudKit.CKDatabaseOperation : CloudKit.CKOperation
*/

type CKDatabaseOperation struct {
  *CKOperation

}

func (r *CKDatabaseOperation) SetDatabase() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKDatabaseOperation) Database() (*CKDatabase, error) {
  return nil, fmt.Errorf("unimplemented")
}

