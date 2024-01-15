//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_INQUIRY_PageC0_Data
*/

type SCSICmdINQUIRYPageC0Data struct {
  FHeader SCSICmdINQUIRYPAGECxHeader `v8:"fHeader"`
  FTdmPageVersion byte `v8:"fTdmPageVersion"`
  FTdmProtocolVersion byte `v8:"fTdmProtocolVersion"`
  FReserved1 byte `v8:"fReserved1"`
  FReserved2 byte `v8:"fReserved2"`
  FMacModelId [32]byte `v8:"fMacModelId"`
  FSerialNumber [32]byte `v8:"fSerialNumber"`
  FMaxReadSize uint `v8:"fMaxReadSize"`
  FMaxWriteSize uint `v8:"fMaxWriteSize"`
  FNativeBlockSize uint `v8:"fNativeBlockSize"`
  FPreferredIOSize uint `v8:"fPreferredIOSize"`
  FFeatures uint64 `v8:"fFeatures"`
  FWorkArounds uint64 `v8:"fWorkArounds"`
  FEncryptionType uint16 `v8:"fEncryptionType"`
  FReserved3 [2]byte `v8:"fReserved3"`
  FInstalledRAMSize uint64 `v8:"fInstalledRAMSize"`
}
