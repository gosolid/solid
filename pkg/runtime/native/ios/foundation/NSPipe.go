//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSPipe : objc.NSObject
*/

type NSPipe struct {
  *objc.NSObject

}

func (r *NSPipe) Pipe() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPipe) FileHandleForReading() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPipe) FileHandleForWriting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

