//js:package native/macos/foundation
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

func (r *NSOutputStream) InitWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOutputStream) HasSpaceAvailable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOutputStream) Write() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOutputStream) InitToMemory() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOutputStream) InitToBuffer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

