//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSURLSessionTaskMetrics : objc.NSObject
*/

type NSURLSessionTaskMetrics struct {
  *objc.NSObject

}

func (r *NSURLSessionTaskMetrics) TaskInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskMetrics) RedirectCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskMetrics) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskMetrics) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskMetrics) TransactionMetrics() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

