//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UserNotifications.UNNotificationSound : objc.NSObject
*/

type UNNotificationSound struct {
  *objc.NSObject

}

func (r *UNNotificationSound) DefaultCriticalSound() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSound) DefaultCriticalSoundWithAudioVolume() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSound) SoundNamed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSound) RingtoneSoundNamed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSound) CriticalSoundNamed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSound) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSound) DefaultSound() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSound) DefaultRingtoneSound() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

