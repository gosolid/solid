//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Metal.MTLStructType : Metal.MTLType
*/

type MTLStructType struct {
  *MTLType

}

func (r *MTLStructType) MemberByName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStructType) Members() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

