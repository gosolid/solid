//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_INQUIRY_PageB2_Data
*/

type SCSICmdINQUIRYPageB2Data struct {
  PERIPHERALDEVICETYPE byte `v8:"pERIPHERALDEVICETYPE"`
  PAGECODE byte `v8:"pAGECODE"`
  PAGELENGTH uint16 `v8:"pAGELENGTH"`
  THRESHOLDEXPONENT byte `v8:"tHRESHOLDEXPONENT"`
  LBPFLAGS byte `v8:"lBPFLAGS"`
  MINIMUMPERCENTAGE byte `v8:"mINIMUMPERCENTAGE"`
  THRESHOLDPERCENTAGE byte `v8:"tHRESHOLDPERCENTAGE"`
}
