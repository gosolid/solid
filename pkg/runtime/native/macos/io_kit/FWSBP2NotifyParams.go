//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.FWSBP2NotifyParams
*/

type FWSBP2NotifyParams struct {
  RefCon void `v8:"refCon"`
  NotificationEvent uint `v8:"notificationEvent"`
  Message void `v8:"message"`
  Length uint `v8:"length"`
  Generation uint `v8:"generation"`
}
