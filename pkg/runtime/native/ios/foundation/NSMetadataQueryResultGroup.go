//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSMetadataQueryResultGroup : objc.NSObject
*/

type NSMetadataQueryResultGroup struct {
  *objc.NSObject

}

func (r *NSMetadataQueryResultGroup) Value() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQueryResultGroup) Subgroups() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQueryResultGroup) ResultCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQueryResultGroup) Results() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQueryResultGroup) ResultAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQueryResultGroup) Attribute() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

