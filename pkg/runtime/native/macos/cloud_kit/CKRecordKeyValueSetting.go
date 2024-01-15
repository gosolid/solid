//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol CloudKit.CKRecordKeyValueSetting
*/

type CKRecordKeyValueSetting struct {

}

func (r *CKRecordKeyValueSetting) AllKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordKeyValueSetting) ChangedKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordKeyValueSetting) ObjectForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordKeyValueSetting) SetObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKRecordKeyValueSetting) ObjectForKeyedSubscript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

