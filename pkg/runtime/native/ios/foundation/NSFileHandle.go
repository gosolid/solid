//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSFileHandle : objc.NSObject
*/

type NSFileHandle struct {
  *objc.NSObject

}

func (r *NSFileHandle) GetOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) SeekToOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) ReadDataUpToLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) WriteData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) TruncateAtOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) SynchronizeAndReturnError() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) CloseAndReturnError() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) AvailableData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) InitWithFileDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) ReadDataToEndOfFileAndReturnError() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileHandle) SeekToEndReturningOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

