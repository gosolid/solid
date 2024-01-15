//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDMultiConnectInfoRec
*/

type VDMultiConnectInfoRec struct {
  CsDisplayCountOrNumber uint `v8:"csDisplayCountOrNumber"`
  CsConnectInfo VDDisplayConnectInfoRec `v8:"csConnectInfo"`
}
