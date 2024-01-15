//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_INQUIRY_PageB1_Data
*/

type SCSICmdINQUIRYPageB1Data struct {
  PERIPHERALDEVICETYPE byte `v8:"pERIPHERALDEVICETYPE"`
  PAGECODE byte `v8:"pAGECODE"`
  Reserved byte `v8:"reserved"`
  PAGELENGTH byte `v8:"pAGELENGTH"`
  MEDIUMROTATIONRATE uint16 `v8:"mEDIUMROTATIONRATE"`
  Reserved2 [58]byte `v8:"reserved2"`
}
