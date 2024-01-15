//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSMutableFontCollection : AppKit.NSFontCollection
*/

type NSMutableFontCollection struct {
  *NSFontCollection

}

func (r *NSMutableFontCollection) FontCollectionWithName() (*NSMutableFontCollection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableFontCollection) RemoveQueryForDescriptors() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableFontCollection) SetQueryDescriptors() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableFontCollection) SetExclusionDescriptors() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableFontCollection) FontCollectionWithDescriptors() (*NSMutableFontCollection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableFontCollection) AddQueryForDescriptors() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableFontCollection) FontCollectionWithAllAvailableDescriptors() (*NSMutableFontCollection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableFontCollection) QueryDescriptors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableFontCollection) ExclusionDescriptors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableFontCollection) FontCollectionWithLocale() (*NSMutableFontCollection, error) {
  return nil, fmt.Errorf("unimplemented")
}

