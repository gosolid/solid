//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLSession : objc.NSObject
*/

type NSURLSession struct {
  *objc.NSObject

}

func (r *NSURLSession) FinishTasksAndInvalidate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSession) UploadTaskWithResumeData() (*NSURLSessionUploadTask, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) WebSocketTaskWithRequest() (*NSURLSessionWebSocketTask, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) FlushWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSession) DownloadTaskWithRequest() (*NSURLSessionDownloadTask, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) WebSocketTaskWithURL() (*NSURLSessionWebSocketTask, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) SetSessionDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSession) GetTasksWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSession) SessionWithConfiguration() (*NSURLSession, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) ResetWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSession) DataTaskWithRequest() (*NSURLSessionDataTask, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) DownloadTaskWithURL() (*NSURLSessionDownloadTask, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) DownloadTaskWithResumeData() (*NSURLSessionDownloadTask, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) SessionDescription() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) UploadTaskWithRequest() (*NSURLSessionUploadTask, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) DelegateQueue() (*NSOperationQueue, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) DataTaskWithURL() (*NSURLSessionDataTask, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) StreamTaskWithNetService() (*NSURLSessionStreamTask, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) GetAllTasksWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSession) StreamTaskWithHostName() (*NSURLSessionStreamTask, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) Configuration() (*NSURLSessionConfiguration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) InvalidateAndCancel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSession) UploadTaskWithStreamedRequest() (*NSURLSessionUploadTask, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSession) SharedSession() (*NSURLSession, error) {
  return nil, fmt.Errorf("unimplemented")
}

