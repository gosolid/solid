//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSItemProvider : objc.NSObject
*/

type NSItemProvider struct {
  *objc.NSObject
  *NSCopying
}

func (r *NSItemProvider) RegisteredTypeIdentifiers() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) SuggestedName() (*objc.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) SetSuggestedName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) HasRepresentationConformingToTypeIdentifier() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) LoadDataRepresentationForTypeIdentifier() (*NSProgress, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) RegisterObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) CanLoadObjectOfClass() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) LoadInPlaceFileRepresentationForTypeIdentifier() (*NSProgress, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) InitWithContentsOfURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) RegisterItemForTypeIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) LoadItemForTypeIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) RegisterDataRepresentationForTypeIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) RegisterFileRepresentationForTypeIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) HasItemConformingToTypeIdentifier() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) RegisterObjectOfClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) InitWithItem() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) RegisteredTypeIdentifiersWithFileOptions() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) LoadFileRepresentationForTypeIdentifier() (*NSProgress, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) InitWithObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProvider) LoadObjectOfClass() (*NSProgress, error) {
  return nil, fmt.Errorf("unimplemented")
}

