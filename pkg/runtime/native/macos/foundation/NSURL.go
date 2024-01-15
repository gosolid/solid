//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURL : objc.NSObject
*/

type NSURL struct {
  *objc.NSObject
  *NSSecureCoding
  *NSCopying
}

func (r *NSURL) InitWithScheme() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) InitWithDataRepresentation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) StopAccessingSecurityScopedResource() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURL) URLWithDataRepresentation() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) GetFileSystemRepresentation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURL) SetTemporaryResourceValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURL) User() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) FileURLWithFileSystemRepresentation() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) InitByResolvingBookmarkData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) SetResourceValues() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURL) WriteBookmarkData() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURL) BaseURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) ResourceSpecifier() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) FilePathURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) InitWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) CheckResourceIsReachableAndReturnError() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURL) RemoveAllCachedResourceValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURL) URLByResolvingAliasFileAtURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) HasDirectoryPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURL) BookmarkDataWithOptions() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) URLByResolvingBookmarkData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) Path() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) Fragment() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) ParameterString() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) InitFileURLWithFileSystemRepresentation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) ResourceValuesForKeys() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) RelativePath() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) Password() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) Query() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) IsFileURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURL) AbsoluteURLWithDataRepresentation() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) SetResourceValue() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURL) RemoveCachedResourceValueForKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURL) AbsoluteString() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) AbsoluteURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) FileReferenceURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) FileSystemRepresentation() (byte, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURL) URLWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) Host() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) GetResourceValue() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURL) BookmarkDataWithContentsOfURL() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) StartAccessingSecurityScopedResource() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURL) FileURLWithPath() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) InitAbsoluteURLWithDataRepresentation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) DataRepresentation() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) StandardizedURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) InitFileURLWithPath() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) RelativeString() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) Scheme() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) IsFileReferenceURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURL) Port() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

