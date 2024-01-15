//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.ModePageFormatHeader
*/

type ModePageFormatHeader struct {
  PSPAGECODE byte `v8:"pSPAGECODE"`
  PAGELENGTH byte `v8:"pAGELENGTH"`
}
