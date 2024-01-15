//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLStructType : Metal.MTLType
*/

type MTLStructType struct {
  *MTLType

}

func (r *MTLStructType) MemberByName() (*MTLStructMember, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStructType) Members() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

