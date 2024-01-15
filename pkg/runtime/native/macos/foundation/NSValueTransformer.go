//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSValueTransformer : objc.NSObject
*/

type NSValueTransformer struct {
  *objc.NSObject

}

func (r *NSValueTransformer) TransformedValueClass() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSValueTransformer) AllowsReverseTransformation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSValueTransformer) TransformedValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSValueTransformer) ReverseTransformedValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSValueTransformer) SetValueTransformer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSValueTransformer) ValueTransformerForName() (*NSValueTransformer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSValueTransformer) ValueTransformerNames() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

