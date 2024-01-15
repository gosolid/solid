//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSSharingService : objc.NSObject
*/

type NSSharingService struct {
  *objc.NSObject

}

func (r *NSSharingService) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) Subject() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) MessageBody() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) PermanentLink() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) AttachmentFileURLs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) CanPerformWithItems() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) PerformWithItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSharingService) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) AlternateImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) SetMenuItemTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSharingService) SetRecipients() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSharingService) SharingServicesForItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) SharingServiceNamed() (*NSSharingService, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) MenuItemTitle() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) Recipients() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) SetSubject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSharingService) AccountName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) InitWithTitle() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingService) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

