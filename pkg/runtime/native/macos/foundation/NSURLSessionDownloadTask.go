//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSURLSessionDownloadTask : Foundation.NSURLSessionTask
*/

type NSURLSessionDownloadTask struct {
  *NSURLSessionTask

}

func (r *NSURLSessionDownloadTask) CancelByProducingResumeData() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionDownloadTask) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionDownloadTask) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

