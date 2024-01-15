//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKNotificationInfo : objc.NSObject
*/

type CKNotificationInfo struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKNotificationInfo) AlertActionLocalizationKey() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SoundName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetShouldSendContentAvailable() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetAlertLocalizationKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetTitleLocalizationArgs() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SubtitleLocalizationArgs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) ShouldBadge() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) ShouldSendMutableContent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetShouldSendMutableContent() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetSoundName() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) DesiredKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) AlertBody() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) AlertLocalizationArgs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) TitleLocalizationKey() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) TitleLocalizationArgs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetSubtitleLocalizationKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetSubtitleLocalizationArgs() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) CollapseIDKey() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) Subtitle() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetAlertActionLocalizationKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetShouldBadge() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) ShouldSendContentAvailable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) Category() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SubtitleLocalizationKey() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) AlertLaunchImage() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetAlertLaunchImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) AlertLocalizationKey() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetAlertLocalizationArgs() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetTitleLocalizationKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetSubtitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetCategory() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetCollapseIDKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetAlertBody() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKNotificationInfo) SetDesiredKeys() error {
  return fmt.Errorf("unimplemented")
}

