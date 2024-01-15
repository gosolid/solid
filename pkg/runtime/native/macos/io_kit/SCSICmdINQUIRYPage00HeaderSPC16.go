//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_INQUIRY_Page00_Header_SPC_16
*/

type SCSICmdINQUIRYPage00HeaderSPC16 struct {
  PERIPHERALDEVICETYPE byte `v8:"pERIPHERALDEVICETYPE"`
  PAGECODE byte `v8:"pAGECODE"`
  PAGELENGTH uint16 `v8:"pAGELENGTH"`
}
