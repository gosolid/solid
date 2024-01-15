//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSInvocationOperation : Foundation.NSOperation
*/

type NSInvocationOperation struct {
  *NSOperation

}

func (r *NSInvocationOperation) InitWithInvocation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocationOperation) Invocation() (*NSInvocation, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocationOperation) Result() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocationOperation) InitWithTarget() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

