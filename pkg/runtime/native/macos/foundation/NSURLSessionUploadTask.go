//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSURLSessionUploadTask : Foundation.NSURLSessionDataTask
*/

type NSURLSessionUploadTask struct {
  *NSURLSessionDataTask

}

func (r *NSURLSessionUploadTask) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionUploadTask) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionUploadTask) CancelByProducingResumeData() error {
  return fmt.Errorf("unimplemented")
}

