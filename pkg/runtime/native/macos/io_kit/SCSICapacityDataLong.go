//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSI_Capacity_Data_Long
*/

type SCSICapacityDataLong struct {
  RETURNEDLOGICALBLOCKADDRESS uint64 `v8:"rETURNEDLOGICALBLOCKADDRESS"`
  BLOCKLENGTHINBYTES uint `v8:"bLOCKLENGTHINBYTES"`
  RTOENPROTEN byte `v8:"rTOENPROTEN"`
  Reserved [19]byte `v8:"reserved"`
}
