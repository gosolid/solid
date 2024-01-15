//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_INQUIRY_Page80_Header
*/

type SCSICmdINQUIRYPage80Header struct {
  PERIPHERALDEVICETYPE byte `v8:"pERIPHERALDEVICETYPE"`
  PAGECODE byte `v8:"pAGECODE"`
  RESERVED byte `v8:"rESERVED"`
  PAGELENGTH byte `v8:"pAGELENGTH"`
  PRODUCTSERIALNUMBER byte `v8:"pRODUCTSERIALNUMBER"`
}
