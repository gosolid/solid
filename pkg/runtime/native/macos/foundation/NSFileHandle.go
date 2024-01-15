//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSFileHandle : objc.NSObject
*/

type NSFileHandle struct {
  *objc.NSObject
  *NSSecureCoding
}

func (r *NSFileHandle) SeekToEndReturningOffset() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) TruncateAtOffset() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) SynchronizeAndReturnError() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) CloseAndReturnError() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) AvailableData() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) InitWithFileDescriptor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) ReadDataToEndOfFileAndReturnError() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) GetOffset() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) SeekToOffset() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) ReadDataUpToLength() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) WriteData() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

