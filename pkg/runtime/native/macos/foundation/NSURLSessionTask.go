//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLSessionTask : objc.NSObject
*/

type NSURLSessionTask struct {
  *objc.NSObject
  *NSCopying
  *NSProgressReporting
}

func (r *NSURLSessionTask) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) PrefersIncrementalDelivery() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Cancel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) OriginalRequest() (*NSURLRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) SetEarliestBeginDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) State() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) SetPriority() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Response() (*NSURLResponse, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) CountOfBytesClientExpectsToReceive() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Error() (*NSError, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Progress() (*NSProgress, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) EarliestBeginDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) SetCountOfBytesClientExpectsToSend() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) CountOfBytesReceived() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Resume() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) TaskIdentifier() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) SetCountOfBytesClientExpectsToReceive() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) CountOfBytesExpectedToSend() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) CountOfBytesExpectedToReceive() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Priority() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) CurrentRequest() (*NSURLRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) SetPrefersIncrementalDelivery() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) CountOfBytesClientExpectsToSend() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) CountOfBytesSent() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) SetTaskDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) Suspend() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTask) TaskDescription() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

