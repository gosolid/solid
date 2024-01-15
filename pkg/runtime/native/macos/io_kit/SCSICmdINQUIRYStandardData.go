//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_INQUIRY_StandardData
*/

type SCSICmdINQUIRYStandardData struct {
  PERIPHERALDEVICETYPE byte `v8:"pERIPHERALDEVICETYPE"`
  RMB byte `v8:"rMB"`
  VERSION byte `v8:"vERSION"`
  RESPONSEDATAFORMAT byte `v8:"rESPONSEDATAFORMAT"`
  ADDITIONALLENGTH byte `v8:"aDDITIONALLENGTH"`
  SCCSReserved byte `v8:"sCCSReserved"`
  Flags1 byte `v8:"flags1"`
  Flags2 byte `v8:"flags2"`
  VENDORIDENTIFICATION [8]byte `v8:"vENDORIDENTIFICATION"`
  PRODUCTIDENTIFICATION [16]byte `v8:"pRODUCTIDENTIFICATION"`
  PRODUCTREVISIONLEVEL [4]byte `v8:"pRODUCTREVISIONLEVEL"`
}
