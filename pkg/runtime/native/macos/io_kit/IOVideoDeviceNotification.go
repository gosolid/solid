//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOVideoDeviceNotification
*/

type IOVideoDeviceNotification struct {
  MObjectID uint `v8:"mObjectID"`
  MNotificationID uint `v8:"mNotificationID"`
  MNotificationArgument1 uint `v8:"mNotificationArgument1"`
  MNotificationArgument2 uint `v8:"mNotificationArgument2"`
  MNotificationArgument3 uint64 `v8:"mNotificationArgument3"`
  MNotificationArgument4 uint64 `v8:"mNotificationArgument4"`
}
