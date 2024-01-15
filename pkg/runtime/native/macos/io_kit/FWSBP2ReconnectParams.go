//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.FWSBP2ReconnectParams
*/

type FWSBP2ReconnectParams struct {
  RefCon void `v8:"refCon"`
  Generation uint `v8:"generation"`
  Status int `v8:"status"`
  ReconnectStatusBlock FWSBP2StatusBlock `v8:"reconnectStatusBlock"`
  ReconnectStatusBlockLength uint `v8:"reconnectStatusBlockLength"`
}
