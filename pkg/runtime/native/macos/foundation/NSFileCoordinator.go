//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSFileCoordinator : objc.NSObject
*/

type NSFileCoordinator struct {
  *objc.NSObject

}

func (r *NSFileCoordinator) ItemAtURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) Cancel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) FilePresenters() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) CoordinateAccessWithIntents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) PrepareForReadingItemsAtURLs() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) InitWithFilePresenter() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) CoordinateReadingItemAtURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) CoordinateWritingItemAtURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) PurposeIdentifier() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) SetPurposeIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) AddFilePresenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) RemoveFilePresenter() error {
  return fmt.Errorf("unimplemented")
}

