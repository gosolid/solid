//js:package native/macos/foundation
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

func (r *NSURLSessionStreamTask) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) CloseWrite() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) CloseRead() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) StartSecureConnection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) StopSecureConnection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) ReadDataOfMinLength() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) WriteData() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionStreamTask) CaptureStreams() error {
  return fmt.Errorf("unimplemented")
}

