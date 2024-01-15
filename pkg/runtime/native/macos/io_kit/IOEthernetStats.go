//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOEthernetStats
*/

type IOEthernetStats struct {
  Dot3StatsEntry IODot3StatsEntry `v8:"dot3StatsEntry"`
  Dot3CollEntry IODot3CollEntry `v8:"dot3CollEntry"`
  Dot3RxExtraEntry IODot3RxExtraEntry `v8:"dot3RxExtraEntry"`
  Dot3TxExtraEntry IODot3TxExtraEntry `v8:"dot3TxExtraEntry"`
}
