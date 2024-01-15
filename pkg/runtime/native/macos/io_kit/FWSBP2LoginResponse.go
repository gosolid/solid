//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.FWSBP2LoginResponse
*/

type FWSBP2LoginResponse struct {
  Length uint16 `v8:"length"`
  LoginID uint16 `v8:"loginID"`
  CommandBlockAgentAddressHi uint `v8:"commandBlockAgentAddressHi"`
  CommandBlockAgentAddressLo uint `v8:"commandBlockAgentAddressLo"`
  Reserved uint16 `v8:"reserved"`
  ReconnectHold uint16 `v8:"reconnectHold"`
}
