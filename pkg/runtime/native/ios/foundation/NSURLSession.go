//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSURLSession : objc.NSObject
*/

type NSURLSession struct {
  *objc.NSObject

}

func (r *NSURLSession) GetTasksWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) StreamTaskWithNetService() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) InvalidateAndCancel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) FlushWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) SharedSession() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) SessionDescription() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) Configuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) DelegateQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) SessionWithConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) DownloadTaskWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) DataTaskWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) UploadTaskWithStreamedRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) StreamTaskWithHostName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) WebSocketTaskWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) WebSocketTaskWithRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) SetSessionDescription() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) GetAllTasksWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) DataTaskWithRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) ResetWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) DownloadTaskWithRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) UploadTaskWithResumeData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) DownloadTaskWithResumeData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) FinishTasksAndInvalidate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) UploadTaskWithRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

