//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSDistantObjectRequest : objc.NSObject
*/

type NSDistantObjectRequest struct {
  *objc.NSObject

}

func (r *NSDistantObjectRequest) ReplyWithException() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDistantObjectRequest) Invocation() (*NSInvocation, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDistantObjectRequest) Connection() (*NSConnection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDistantObjectRequest) Conversation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

