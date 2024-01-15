//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSURLSessionStreamTask : Foundation.NSURLSessionTask
*/

type NSURLSessionStreamTask struct {
  *NSURLSessionTask

}

func (r *NSURLSessionStreamTask) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) CaptureStreams() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) CloseRead() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) CloseWrite() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) StartSecureConnection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) StopSecureConnection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) ReadDataOfMinLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) WriteData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

