//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.CDATIP
*/

type CDATIP struct {
  DataLength uint16 `v8:"dataLength"`
  Reserved [2]byte `v8:"reserved"`
  ReferenceSpeed byte `v8:"referenceSpeed"`
  Reserved3 byte `v8:"reserved3"`
  IndicativeTargetWritingPower byte `v8:"indicativeTargetWritingPower"`
  Reserved2 byte `v8:"reserved2"`
  Reserved5 byte `v8:"reserved5"`
  UnrestrictedUse byte `v8:"unrestrictedUse"`
  Reserved4 byte `v8:"reserved4"`
  A3Valid byte `v8:"a3Valid"`
  A2Valid byte `v8:"a2Valid"`
  A1Valid byte `v8:"a1Valid"`
  DiscSubType byte `v8:"discSubType"`
  DiscType byte `v8:"discType"`
  Reserved6 byte `v8:"reserved6"`
  Reserved7 byte `v8:"reserved7"`
  StartTimeOfLeadIn CDMSF `v8:"startTimeOfLeadIn"`
  Reserved8 byte `v8:"reserved8"`
  LastPossibleStartTimeOfLeadOut CDMSF `v8:"lastPossibleStartTimeOfLeadOut"`
  Reserved9 byte `v8:"reserved9"`
  A1 [3]byte `v8:"a1"`
  Reserved10 byte `v8:"reserved10"`
  A2 [3]byte `v8:"a2"`
  Reserved11 byte `v8:"reserved11"`
  A3 [3]byte `v8:"a3"`
  Reserved12 byte `v8:"reserved12"`
}
