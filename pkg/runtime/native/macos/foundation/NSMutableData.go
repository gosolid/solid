//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMutableData : Foundation.NSData
*/

type NSMutableData struct {
  *NSData

}

func (r *NSMutableData) MutableBytes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableData) Length() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableData) SetLength() error {
  return fmt.Errorf("unimplemented")
}

