//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSFileVersion : objc.NSObject
*/

type NSFileVersion struct {
  *objc.NSObject

}

func (r *NSFileVersion) GetNonlocalVersionsOfItemAtURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) TemporaryDirectoryURLForNewVersionOfItemAtURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) ReplaceItemAtURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) IsConflict() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) SetResolved() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) SetDiscardable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) HasLocalContents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) HasThumbnail() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) OtherVersionsOfItemAtURL() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) UnresolvedConflictVersionsOfItemAtURL() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) LocalizedName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) OriginatorNameComponents() (*NSPersonNameComponents, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) CurrentVersionOfItemAtURL() (*NSFileVersion, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) RemoveAndReturnError() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) RemoveOtherVersionsOfItemAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) LocalizedNameOfSavingComputer() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) ModificationDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) IsDiscardable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) VersionOfItemAtURL() (*NSFileVersion, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) AddVersionOfItemAtURL() (*NSFileVersion, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) URL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) PersistentIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileVersion) IsResolved() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

