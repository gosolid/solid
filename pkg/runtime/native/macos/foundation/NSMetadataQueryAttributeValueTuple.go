//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSMetadataQueryAttributeValueTuple : objc.NSObject
*/

type NSMetadataQueryAttributeValueTuple struct {
  *objc.NSObject

}

func (r *NSMetadataQueryAttributeValueTuple) Attribute() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQueryAttributeValueTuple) Value() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQueryAttributeValueTuple) Count() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

