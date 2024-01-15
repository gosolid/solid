//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSMetadataItem : objc.NSObject
*/

type NSMetadataItem struct {
  *objc.NSObject

}

func (r *NSMetadataItem) Attributes() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataItem) InitWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataItem) ValueForAttribute() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataItem) ValuesForAttributes() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

