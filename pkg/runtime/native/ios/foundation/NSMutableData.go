//js:package native/ios/foundation
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

func (r *NSMutableData) Length() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableData) SetLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableData) MutableBytes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

