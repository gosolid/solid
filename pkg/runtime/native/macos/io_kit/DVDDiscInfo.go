//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DVDDiscInfo
*/

type DVDDiscInfo struct {
  DataLength uint16 `v8:"dataLength"`
  DiscStatus byte `v8:"discStatus"`
  StateOfLastBorder byte `v8:"stateOfLastBorder"`
  Erasable byte `v8:"erasable"`
  Reserved byte `v8:"reserved"`
  Reserved2 byte `v8:"reserved2"`
  NumberOfBordersLSB byte `v8:"numberOfBordersLSB"`
  FirstRZoneNumberInLastBorderLSB byte `v8:"firstRZoneNumberInLastBorderLSB"`
  LastRZoneNumberInLastBorderLSB byte `v8:"lastRZoneNumberInLastBorderLSB"`
  Reserved3 byte `v8:"reserved3"`
  UnrestrictedUse byte `v8:"unrestrictedUse"`
  DiscBarCodeValid byte `v8:"discBarCodeValid"`
  Reserved4 byte `v8:"reserved4"`
  Reserved5 byte `v8:"reserved5"`
  NumberOfBordersMSB byte `v8:"numberOfBordersMSB"`
  FirstRZoneNumberInLastBorderMSB byte `v8:"firstRZoneNumberInLastBorderMSB"`
  LastRZoneNumberInLastBorderMSB byte `v8:"lastRZoneNumberInLastBorderMSB"`
  Reserved6 [4]byte `v8:"reserved6"`
  Reserved7 [4]byte `v8:"reserved7"`
  Reserved8 [4]byte `v8:"reserved8"`
  DiscBarCode [8]byte `v8:"discBarCode"`
  Reserved9 byte `v8:"reserved9"`
  NumberOfOPCTableEntries byte `v8:"numberOfOPCTableEntries"`
  OpcTableEntries [0]byte `v8:"opcTableEntries"`
}
