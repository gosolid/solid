//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UILocalNotification : objc.NSObject
*/

type UILocalNotification struct {
  *objc.NSObject

}

func (r *UILocalNotification) SetRepeatCalendar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) Region() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) SetAlertLaunchImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) RepeatCalendar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) SetRegion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) AlertBody() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) HasAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) SetHasAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) AlertTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) SoundName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) ApplicationIconBadgeNumber() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) SetFireDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) SetRepeatInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) SetRegionTriggersOnce() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) SetAlertTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) SetUserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) RepeatInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) RegionTriggersOnce() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) SetAlertBody() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) AlertLaunchImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) SetSoundName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) FireDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) SetCategory() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) AlertAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) SetTimeZone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) Category() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) TimeZone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) SetApplicationIconBadgeNumber() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalNotification) SetAlertAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

