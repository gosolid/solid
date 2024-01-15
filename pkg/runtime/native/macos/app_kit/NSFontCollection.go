//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSFontCollection : objc.NSObject
*/

type NSFontCollection struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSMutableCopying
  *foundation.NSCoding
}

func (r *NSFontCollection) HideFontCollectionWithName() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFontCollection) FontCollectionWithAllAvailableDescriptors() (*NSFontCollection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontCollection) QueryDescriptors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontCollection) ExclusionDescriptors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontCollection) FontCollectionWithDescriptors() (*NSFontCollection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontCollection) RenameFontCollectionWithName() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFontCollection) MatchingDescriptorsForFamily() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontCollection) AllFontCollectionNames() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontCollection) ShowFontCollection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFontCollection) FontCollectionWithName() (*NSFontCollection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontCollection) FontCollectionWithLocale() (*NSFontCollection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontCollection) MatchingDescriptorsWithOptions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontCollection) MatchingDescriptors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

