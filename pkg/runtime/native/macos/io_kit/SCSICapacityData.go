//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSI_Capacity_Data
*/

type SCSICapacityData struct {
  RETURNEDLOGICALBLOCKADDRESS uint `v8:"rETURNEDLOGICALBLOCKADDRESS"`
  BLOCKLENGTHINBYTES uint `v8:"bLOCKLENGTHINBYTES"`
}
