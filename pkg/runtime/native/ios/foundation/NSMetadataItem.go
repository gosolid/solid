//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSMetadataItem : objc.NSObject
*/

type NSMetadataItem struct {
  *objc.NSObject

}

func (r *NSMetadataItem) InitWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataItem) ValueForAttribute() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataItem) ValuesForAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataItem) Attributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

