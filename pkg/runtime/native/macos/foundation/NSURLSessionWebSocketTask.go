//js:package native/macos/foundation
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

func (r *NSURLSessionWebSocketTask) CancelWithCloseCode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) SetMaximumMessageSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) CloseCode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) SendMessage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) SendPingWithPongReceiveHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) CloseReason() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) ReceiveMessageWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionWebSocketTask) MaximumMessageSize() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

