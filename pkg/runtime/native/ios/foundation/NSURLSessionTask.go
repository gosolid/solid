//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSURLSessionTask : objc.NSObject
*/

type NSURLSessionTask struct {
  *objc.NSObject

}

func (r *NSURLSessionTask) Priority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) EarliestBeginDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) CountOfBytesExpectedToSend() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Error() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Progress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) CountOfBytesReceived() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) CountOfBytesClientExpectsToSend() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) TaskDescription() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Cancel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) CurrentRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Response() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) SetPriority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Resume() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) SetCountOfBytesClientExpectsToSend() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) CountOfBytesSent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) SetPrefersIncrementalDelivery() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Suspend() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) TaskIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) CountOfBytesClientExpectsToReceive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) SetTaskDescription() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) CountOfBytesExpectedToReceive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) State() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) PrefersIncrementalDelivery() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) OriginalRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) SetEarliestBeginDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) SetCountOfBytesClientExpectsToReceive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

