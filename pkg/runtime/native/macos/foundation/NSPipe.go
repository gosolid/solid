//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSPipe : objc.NSObject
*/

type NSPipe struct {
  *objc.NSObject

}

func (r *NSPipe) Pipe() (*NSPipe, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPipe) FileHandleForReading() (*NSFileHandle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPipe) FileHandleForWriting() (*NSFileHandle, error) {
  return nil, fmt.Errorf("unimplemented")
}

