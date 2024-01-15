//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSURLSessionWebSocketTask : Foundation.NSURLSessionTask
*/

type NSURLSessionWebSocketTask struct {
  *NSURLSessionTask

}

func (r *NSURLSessionWebSocketTask) SendMessage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) CancelWithCloseCode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) CloseCode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) CloseReason() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) MaximumMessageSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) SetMaximumMessageSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) ReceiveMessageWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) SendPingWithPongReceiveHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

