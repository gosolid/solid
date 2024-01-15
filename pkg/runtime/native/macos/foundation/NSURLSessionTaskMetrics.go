//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSURLSessionTaskMetrics : objc.NSObject
*/

type NSURLSessionTaskMetrics struct {
  *objc.NSObject

}

func (r *NSURLSessionTaskMetrics) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskMetrics) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskMetrics) TransactionMetrics() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskMetrics) TaskInterval() (*NSDateInterval, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskMetrics) RedirectCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

