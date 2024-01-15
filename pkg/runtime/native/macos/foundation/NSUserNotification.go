//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSUserNotification : objc.NSObject
*/

type NSUserNotification struct {
  *objc.NSObject
  *NSCopying
}

func (r *NSUserNotification) IsPresented() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetSoundName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetContentImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetInformativeText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetSubtitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetActionButtonTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) DeliveryTimeZone() (*NSTimeZone, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetDeliveryRepeatInterval() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) ActualDeliveryDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) ContentImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) Title() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) AdditionalActions() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) IsRemote() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) HasActionButton() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) OtherButtonTitle() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) HasReplyButton() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetHasReplyButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) ResponsePlaceholder() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) Subtitle() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetAdditionalActions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetResponsePlaceholder() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) ActionButtonTitle() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) DeliveryRepeatInterval() (*NSDateComponents, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) ActivationType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetOtherButtonTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SoundName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetHasActionButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) Identifier() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) AdditionalActivationAction() (*NSUserNotificationAction, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetDeliveryDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetUserInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) DeliveryDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) SetDeliveryTimeZone() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) UserInfo() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) Response() (*NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotification) InformativeText() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

