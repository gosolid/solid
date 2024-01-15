//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_INQUIRY_PageC1_Data
*/

type SCSICmdINQUIRYPageC1Data struct {
  FHeader SCSICmdINQUIRYPAGECxHeader `v8:"fHeader"`
  FTdmPowerRequirementsPageVersion byte `v8:"fTdmPowerRequirementsPageVersion"`
  FReserved1 byte `v8:"fReserved1"`
  FReserved2 uint16 `v8:"fReserved2"`
  FPowerRequired uint `v8:"fPowerRequired"`
}
