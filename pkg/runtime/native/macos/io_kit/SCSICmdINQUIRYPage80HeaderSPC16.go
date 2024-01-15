//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_INQUIRY_Page80_Header_SPC_16
*/

type SCSICmdINQUIRYPage80HeaderSPC16 struct {
  PERIPHERALDEVICETYPE byte `v8:"pERIPHERALDEVICETYPE"`
  PAGECODE byte `v8:"pAGECODE"`
  PAGELENGTH uint16 `v8:"pAGELENGTH"`
  PRODUCTSERIALNUMBER byte `v8:"pRODUCTSERIALNUMBER"`
}
