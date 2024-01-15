//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.FWSBP2LogoutCompleteParams
*/

type FWSBP2LogoutCompleteParams struct {
  RefCon void `v8:"refCon"`
  Generation uint `v8:"generation"`
  Status int `v8:"status"`
  StatusBlock FWSBP2StatusBlock `v8:"statusBlock"`
  StatusBlockLength uint `v8:"statusBlockLength"`
}
