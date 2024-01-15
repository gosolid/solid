//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSFileWrapper : objc.NSObject
*/

type NSFileWrapper struct {
  *objc.NSObject

}

func (r *NSFileWrapper) IsRegularFile() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) RemoveFileWrapper() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) IsSymbolicLink() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) FileWrappers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) RegularFileContents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) AddRegularFileWithContents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) PreferredFilename() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) Filename() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) SetFilename() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) InitWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) InitRegularFileWithContents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) FileAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) SetFileAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) AddFileWrapper() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) IsDirectory() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) InitSymbolicLinkWithDestinationURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) InitWithSerializedRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) WriteToURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) MatchesContentsOfURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) KeyForFileWrapper() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) SerializedRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) InitDirectoryWithFileWrappers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) ReadFromURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) SetPreferredFilename() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileWrapper) SymbolicLinkDestinationURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

