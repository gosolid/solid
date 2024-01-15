//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSBlockOperation : Foundation.NSOperation
*/

type NSBlockOperation struct {
  *NSOperation

}

func (r *NSBlockOperation) BlockOperationWithBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBlockOperation) AddExecutionBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBlockOperation) ExecutionBlocks() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

