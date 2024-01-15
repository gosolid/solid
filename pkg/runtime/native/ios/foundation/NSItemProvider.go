//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSItemProvider : objc.NSObject
*/

type NSItemProvider struct {
  *objc.NSObject

}

func (r *NSItemProvider) HasItemConformingToTypeIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) InitWithObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) RegisteredTypeIdentifiersWithFileOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) RegisterObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) RegisterObjectOfClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) InitWithItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) InitWithContentsOfURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) RegisteredTypeIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) HasRepresentationConformingToTypeIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) LoadDataRepresentationForTypeIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) LoadFileRepresentationForTypeIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) LoadInPlaceFileRepresentationForTypeIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) LoadObjectOfClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) LoadItemForTypeIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) SuggestedName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) SetSuggestedName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) RegisterDataRepresentationForTypeIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) RegisterFileRepresentationForTypeIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) CanLoadObjectOfClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) RegisterItemForTypeIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

