//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIActivityItemsConfiguration : objc.NSObject
*/

type UIActivityItemsConfiguration struct {
  *objc.NSObject

}

func (r *UIActivityItemsConfiguration) LocalObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) SupportedInteractions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) SetMetadataProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) SetPreviewProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) InitWithItemProviders() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) PreviewProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) ApplicationActivitiesProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) SetPerItemMetadataProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) PerItemMetadataProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) SetLocalObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) SetSupportedInteractions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) InitWithObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) MetadataProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) SetApplicationActivitiesProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) ActivityItemsConfigurationWithObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityItemsConfiguration) ActivityItemsConfigurationWithItemProviders() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

