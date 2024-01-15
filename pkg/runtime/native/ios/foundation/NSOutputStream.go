//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSOutputStream : Foundation.NSStream
*/

type NSOutputStream struct {
  *NSStream

}

func (r *NSOutputStream) Write() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOutputStream) InitToMemory() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOutputStream) InitToBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOutputStream) InitWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOutputStream) HasSpaceAvailable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

