//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSFileManager : objc.NSObject
*/

type NSFileManager struct {
  *objc.NSObject

}

func (r *NSFileManager) StringWithFileSystemRepresentation() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ContentsOfDirectoryAtPath() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) AttributesOfItemAtPath() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CopyItemAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) MoveItemAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) RemoveItemAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) MovePath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) IsExecutableFileAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ReplaceItemAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) GetRelationship() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) FileAttributesAtPath() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ChangeFileAttributes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ChangeCurrentDirectoryPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) URLForUbiquityContainerIdentifier() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) UnmountVolumeAtURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CreateDirectoryAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) SubpathsOfDirectoryAtPath() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) LinkItemAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) PathContentOfSymbolicLinkAtPath() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) RemoveFileAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) IsWritableFileAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) IsDeletableFileAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) URLForDirectory() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) AttributesOfFileSystemForPath() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) LinkItemAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) IsReadableFileAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ComponentsToDisplayForPath() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) DestinationOfSymbolicLinkAtPath() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ContentsEqualAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) EnumeratorAtPath() (*NSDirectoryEnumerator, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) SubpathsAtPath() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) SetUbiquitous() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) EvictUbiquitousItemAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) SetAttributes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CreateSymbolicLinkAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CopyItemAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) RemoveItemAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) MountedVolumeURLsIncludingResourceValuesForKeys() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ContentsOfDirectoryAtURL() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) URLForPublishingUbiquitousItemAtURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) DefaultManager() (*NSFileManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) DirectoryContentsAtPath() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) StartDownloadingUbiquitousItemAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) GetFileProviderServicesForItemAtURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileManager) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) UbiquityIdentityToken() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CreateFileAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) FileSystemRepresentationWithPath() (byte, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ContainerURLForSecurityApplicationGroupIdentifier() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) TrashItemAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) FileSystemAttributesAtPath() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) ContentsAtPath() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) MoveItemAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CurrentDirectoryPath() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CreateSymbolicLinkAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) LinkPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CopyPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) CreateDirectoryAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) FileExistsAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) DisplayNameAtPath() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) EnumeratorAtURL() (*NSDirectoryEnumerator, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) IsUbiquitousItemAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileManager) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileManager) URLsForDirectory() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

