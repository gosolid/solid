//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSFileVersion : objc.NSObject
*/

type NSFileVersion struct {
  *objc.NSObject

}

func (r *NSFileVersion) IsResolved() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) RemoveAndReturnError() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) RemoveOtherVersionsOfItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) URL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) OriginatorNameComponents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) ModificationDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) IsConflict() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) PersistentIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) SetResolved() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) SetDiscardable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) CurrentVersionOfItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) VersionOfItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) AddVersionOfItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) LocalizedName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) LocalizedNameOfSavingComputer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) HasThumbnail() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) HasLocalContents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) OtherVersionsOfItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) UnresolvedConflictVersionsOfItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) GetNonlocalVersionsOfItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) TemporaryDirectoryURLForNewVersionOfItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) ReplaceItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) IsDiscardable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

