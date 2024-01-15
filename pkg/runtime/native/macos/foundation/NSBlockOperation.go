//js:package native/macos/foundation
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

func (r *NSBlockOperation) BlockOperationWithBlock() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBlockOperation) AddExecutionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBlockOperation) ExecutionBlocks() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

