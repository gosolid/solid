//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSFileWrapper : objc.NSObject
*/

type NSFileWrapper struct {
  *objc.NSObject
  *NSSecureCoding
}

func (r *NSFileWrapper) InitWithSerializedRepresentation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) IsSymbolicLink() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) SetPreferredFilename() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) FileAttributes() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) FileWrappers() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) RegularFileContents() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) MatchesContentsOfURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) SetFileAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) InitWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) InitSymbolicLinkWithDestinationURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) WriteToURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) AddFileWrapper() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) KeyForFileWrapper() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) AddRegularFileWithContents() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) InitDirectoryWithFileWrappers() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) InitRegularFileWithContents() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) IsDirectory() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) SerializedRepresentation() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) ReadFromURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) SymbolicLinkDestinationURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) RemoveFileWrapper() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) IsRegularFile() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) PreferredFilename() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) Filename() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) SetFilename() error {
  return fmt.Errorf("unimplemented")
}

