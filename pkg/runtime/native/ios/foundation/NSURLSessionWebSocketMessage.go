//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSURLSessionWebSocketMessage : objc.NSObject
*/

type NSURLSessionWebSocketMessage struct {
  *objc.NSObject

}

func (r *NSURLSessionWebSocketMessage) InitWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketMessage) InitWithString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketMessage) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketMessage) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketMessage) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketMessage) Data() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketMessage) String() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

