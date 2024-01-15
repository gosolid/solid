//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSValueTransformer : objc.NSObject
*/

type NSValueTransformer struct {
  *objc.NSObject

}

func (r *NSValueTransformer) ValueTransformerNames() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSValueTransformer) TransformedValueClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSValueTransformer) AllowsReverseTransformation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSValueTransformer) TransformedValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSValueTransformer) ReverseTransformedValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSValueTransformer) SetValueTransformer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSValueTransformer) ValueTransformerForName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

