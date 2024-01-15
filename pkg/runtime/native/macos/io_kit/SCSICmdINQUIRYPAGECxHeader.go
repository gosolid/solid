//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_INQUIRY_PAGECx_Header
*/

type SCSICmdINQUIRYPAGECxHeader struct {
  PERIPHERALDEVICETYPE byte `v8:"pERIPHERALDEVICETYPE"`
  PAGECODE byte `v8:"pAGECODE"`
  RESERVED byte `v8:"rESERVED"`
  PAGELENGTH byte `v8:"pAGELENGTH"`
}
