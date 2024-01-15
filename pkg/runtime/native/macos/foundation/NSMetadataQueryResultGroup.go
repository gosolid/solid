//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSMetadataQueryResultGroup : objc.NSObject
*/

type NSMetadataQueryResultGroup struct {
  *objc.NSObject

}

func (r *NSMetadataQueryResultGroup) Value() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQueryResultGroup) Subgroups() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQueryResultGroup) ResultCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQueryResultGroup) Results() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQueryResultGroup) ResultAtIndex() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQueryResultGroup) Attribute() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

