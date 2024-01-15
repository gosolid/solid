//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSFileManager : objc.NSObject
*/

type NSFileManager struct {
  *objc.NSObject

}

func (r *NSFileManager) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) AttributesOfItemAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) DestinationOfSymbolicLinkAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CopyItemAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) PathContentOfSymbolicLinkAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) FileExistsAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ComponentsToDisplayForPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) SubpathsAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) StartDownloadingUbiquitousItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) SetAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) DisplayNameAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) FileAttributesAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) IsWritableFileAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ContentsEqualAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) FileSystemRepresentationWithPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) UnmountVolumeAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) FileSystemAttributesAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) AttributesOfFileSystemForPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ChangeFileAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) URLForDirectory() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CreateDirectoryAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) TrashItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ChangeCurrentDirectoryPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) UbiquityIdentityToken() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) EvictUbiquitousItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ContainerURLForSecurityApplicationGroupIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) DirectoryContentsAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) IsExecutableFileAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) IsDeletableFileAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) EnumeratorAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ReplaceItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) SetUbiquitous() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CreateDirectoryAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) RemoveItemAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) RemoveItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) IsReadableFileAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) StringWithFileSystemRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) URLForPublishingUbiquitousItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) MountedVolumeURLsIncludingResourceValuesForKeys() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) GetRelationship() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CreateSymbolicLinkAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) LinkItemAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) MoveItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) EnumeratorAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ContentsOfDirectoryAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) MoveItemAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) DefaultManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CreateFileAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) IsUbiquitousItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) URLsForDirectory() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ContentsOfDirectoryAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CreateSymbolicLinkAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CopyItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) LinkItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ContentsAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) SubpathsOfDirectoryAtPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) URLForUbiquityContainerIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) GetFileProviderServicesForItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CurrentDirectoryPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

