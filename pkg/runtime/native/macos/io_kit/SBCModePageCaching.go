//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SBCModePageCaching
*/

type SBCModePageCaching struct {
  Header ModePageFormatHeader `v8:"header"`
  Flags byte `v8:"flags"`
  DEMANDREADWRITERETENTIONPRIORITY byte `v8:"dEMANDREADWRITERETENTIONPRIORITY"`
  DISABLEPREFETCHTRANSFERLENGTH uint16 `v8:"dISABLEPREFETCHTRANSFERLENGTH"`
  MINIMUMPREFETCH uint16 `v8:"mINIMUMPREFETCH"`
  MAXIMUMPREFETCH uint16 `v8:"mAXIMUMPREFETCH"`
  MAXIMUMPREFETCHCEILING uint16 `v8:"mAXIMUMPREFETCHCEILING"`
  Flags2 byte `v8:"flags2"`
  NUMBEROFCACHESEGMENTS byte `v8:"nUMBEROFCACHESEGMENTS"`
  CACHESEGMENTSIZE uint16 `v8:"cACHESEGMENTSIZE"`
  RESERVED byte `v8:"rESERVED"`
  NONCACHESEGMENTSIZE [3]byte `v8:"nONCACHESEGMENTSIZE"`
}
