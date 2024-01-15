//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.CDDiscInfo
*/

type CDDiscInfo struct {
  DataLength uint16 `v8:"dataLength"`
  DiscStatus byte `v8:"discStatus"`
  StateOfLastSession byte `v8:"stateOfLastSession"`
  Erasable byte `v8:"erasable"`
  Reserved byte `v8:"reserved"`
  NumberOfFirstTrack byte `v8:"numberOfFirstTrack"`
  NumberOfSessionsLSB byte `v8:"numberOfSessionsLSB"`
  FirstTrackNumberInLastSessionLSB byte `v8:"firstTrackNumberInLastSessionLSB"`
  LastTrackNumberInLastSessionLSB byte `v8:"lastTrackNumberInLastSessionLSB"`
  Reserved3 byte `v8:"reserved3"`
  UnrestrictedUse byte `v8:"unrestrictedUse"`
  DiscBarCodeValid byte `v8:"discBarCodeValid"`
  DiscIdentificationValid byte `v8:"discIdentificationValid"`
  DiscType byte `v8:"discType"`
  NumberOfSessionsMSB byte `v8:"numberOfSessionsMSB"`
  FirstTrackNumberInLastSessionMSB byte `v8:"firstTrackNumberInLastSessionMSB"`
  LastTrackNumberInLastSessionMSB byte `v8:"lastTrackNumberInLastSessionMSB"`
  DiscIdentification uint `v8:"discIdentification"`
  Reserved7 byte `v8:"reserved7"`
  LastSessionLeadInStartTime CDMSF `v8:"lastSessionLeadInStartTime"`
  Reserved8 byte `v8:"reserved8"`
  LastPossibleStartTimeOfLeadOut CDMSF `v8:"lastPossibleStartTimeOfLeadOut"`
  DiscBarCode [8]byte `v8:"discBarCode"`
  Reserved9 byte `v8:"reserved9"`
  NumberOfOPCTableEntries byte `v8:"numberOfOPCTableEntries"`
  OpcTableEntries [0]byte `v8:"opcTableEntries"`
}
