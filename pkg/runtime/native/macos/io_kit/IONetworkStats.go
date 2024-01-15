//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IONetworkStats
*/

type IONetworkStats struct {
  InputPackets uint `v8:"inputPackets"`
  InputErrors uint `v8:"inputErrors"`
  OutputPackets uint `v8:"outputPackets"`
  OutputErrors uint `v8:"outputErrors"`
  Collisions uint `v8:"collisions"`
}
