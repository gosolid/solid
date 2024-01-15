//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSURL : objc.NSObject
*/

type NSURL struct {
  *objc.NSObject

}

func (r *NSURL) RemoveAllCachedResourceValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) ResourceSpecifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) FileURLWithPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) RemoveCachedResourceValueForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) WriteBookmarkData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) BaseURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) ParameterString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) InitWithScheme() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) InitWithString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) SetResourceValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) StopAccessingSecurityScopedResource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) InitFileURLWithFileSystemRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) AbsoluteURLWithDataRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) AbsoluteString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) FileURLWithFileSystemRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) Scheme() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) Fragment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) FileSystemRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) IsFileURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) InitFileURLWithPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) URLWithDataRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) IsFileReferenceURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) Path() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) StandardizedURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) GetResourceValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) BookmarkDataWithContentsOfURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) URLByResolvingBookmarkData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) SetTemporaryResourceValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) RelativeString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) HasDirectoryPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) InitWithDataRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) SetResourceValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) InitByResolvingBookmarkData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) DataRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) Query() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) CheckResourceIsReachableAndReturnError() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) GetFileSystemRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) URLWithString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) InitAbsoluteURLWithDataRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) URLByResolvingAliasFileAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) AbsoluteURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) User() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) RelativePath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) BookmarkDataWithOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) StartAccessingSecurityScopedResource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) Host() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) Password() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) FileReferenceURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) ResourceValuesForKeys() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) Port() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURL) FilePathURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

