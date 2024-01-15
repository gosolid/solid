//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.BDDiscInfo
*/

type BDDiscInfo struct {
  DataLength uint16 `v8:"dataLength"`
  DiscStatus byte `v8:"discStatus"`
  StateOfLastSession byte `v8:"stateOfLastSession"`
  Erasable byte `v8:"erasable"`
  DataType byte `v8:"dataType"`
  Reserved2 byte `v8:"reserved2"`
  NumberOfSessionsLSB byte `v8:"numberOfSessionsLSB"`
  FirstTrackNumberInLastSessionLSB byte `v8:"firstTrackNumberInLastSessionLSB"`
  LastTrackNumberInLastSessionLSB byte `v8:"lastTrackNumberInLastSessionLSB"`
  Reserved4 [2]byte `v8:"reserved4"`
  NumberOfSessionsMSB byte `v8:"numberOfSessionsMSB"`
  FirstTrackNumberInLastSessionMSB byte `v8:"firstTrackNumberInLastSessionMSB"`
  LastTrackNumberInLastSessionMSB byte `v8:"lastTrackNumberInLastSessionMSB"`
  Reserved6 [22]byte `v8:"reserved6"`
}
