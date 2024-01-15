//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSURLSessionDataTask : Foundation.NSURLSessionTask
*/

type NSURLSessionDataTask struct {
  *NSURLSessionTask

}

func (r *NSURLSessionDataTask) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionDataTask) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

