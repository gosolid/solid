//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSMetadataQueryAttributeValueTuple : objc.NSObject
*/

type NSMetadataQueryAttributeValueTuple struct {
  *objc.NSObject

}

func (r *NSMetadataQueryAttributeValueTuple) Attribute() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQueryAttributeValueTuple) Value() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQueryAttributeValueTuple) Count() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

