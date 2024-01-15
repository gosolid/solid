//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct IOKit.IOVideoDeviceNotificationMessage
*/

type IOVideoDeviceNotificationMessage struct {
  MMessageHeader objc.MachMsgHeaderT `v8:"mMessageHeader"`
  MClientData uint `v8:"mClientData"`
  MNumberNotifications uint `v8:"mNumberNotifications"`
  MNotifications [1]IOVideoDeviceNotification `v8:"mNotifications"`
}
