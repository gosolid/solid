//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DVDTitleKeyInfo
*/

type DVDTitleKeyInfo struct {
  DataLength [2]byte `v8:"dataLength"`
  Reserved [2]byte `v8:"reserved"`
  CPMOD byte `v8:"cPMOD"`
  CGMS byte `v8:"cGMS"`
  CPSEC byte `v8:"cPSEC"`
  CPM byte `v8:"cPM"`
  TitleKeyValue [5]byte `v8:"titleKeyValue"`
  Reserved2 [2]byte `v8:"reserved2"`
}
