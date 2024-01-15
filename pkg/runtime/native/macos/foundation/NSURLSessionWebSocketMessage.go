//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLSessionWebSocketMessage : objc.NSObject
*/

type NSURLSessionWebSocketMessage struct {
  *objc.NSObject

}

func (r *NSURLSessionWebSocketMessage) InitWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketMessage) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketMessage) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketMessage) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketMessage) Data() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketMessage) String() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketMessage) InitWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

