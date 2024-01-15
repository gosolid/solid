//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_INQUIRY_Page89_Data
*/

type SCSICmdINQUIRYPage89Data struct {
  PERIPHERALDEVICETYPE byte `v8:"pERIPHERALDEVICETYPE"`
  PAGECODE byte `v8:"pAGECODE"`
  PAGELENGTH uint16 `v8:"pAGELENGTH"`
  Reserved uint `v8:"reserved"`
  SATVENDORIDENTIFICATION [8]byte `v8:"sATVENDORIDENTIFICATION"`
  SATPRODUCTIDENTIFICATION [16]byte `v8:"sATPRODUCTIDENTIFICATION"`
  SATPRODUCTREVISIONLEVEL [4]byte `v8:"sATPRODUCTREVISIONLEVEL"`
  ATADEVICESIGNATURE [20]byte `v8:"aTADEVICESIGNATURE"`
  COMMANDCODE byte `v8:"cOMMANDCODE"`
  Reserved2 [3]byte `v8:"reserved2"`
  IDENTIFYDATA [512]byte `v8:"iDENTIFYDATA"`
}
