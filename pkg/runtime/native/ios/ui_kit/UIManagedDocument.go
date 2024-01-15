//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIManagedDocument : UIKit.UIDocument
*/

type UIManagedDocument struct {
  *UIDocument

}

func (r *UIManagedDocument) WriteAdditionalContent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIManagedDocument) ManagedObjectContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIManagedDocument) ManagedObjectModel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIManagedDocument) PersistentStoreOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIManagedDocument) SetPersistentStoreOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIManagedDocument) ReadAdditionalContentFromURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIManagedDocument) AdditionalContentForURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIManagedDocument) PersistentStoreName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIManagedDocument) ModelConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIManagedDocument) SetModelConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIManagedDocument) ConfigurePersistentStoreCoordinatorForURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIManagedDocument) PersistentStoreTypeForFileType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

