//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSFileCoordinator : objc.NSObject
*/

type NSFileCoordinator struct {
  *objc.NSObject

}

func (r *NSFileCoordinator) CoordinateReadingItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) PrepareForReadingItemsAtURLs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) ItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) PurposeIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) SetPurposeIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) InitWithFilePresenter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) RemoveFilePresenter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) CoordinateAccessWithIntents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) CoordinateWritingItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) Cancel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) FilePresenters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileCoordinator) AddFilePresenter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

