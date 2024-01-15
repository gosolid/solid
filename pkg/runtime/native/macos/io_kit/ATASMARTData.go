//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.ATASMARTData
*/

type ATASMARTData struct {
  VendorSpecific1 [362]byte `v8:"vendorSpecific1"`
  OffLineDataCollectionStatus byte `v8:"offLineDataCollectionStatus"`
  SelfTestExecutionStatus byte `v8:"selfTestExecutionStatus"`
  SecondsToCompleteOffLineActivity [2]byte `v8:"secondsToCompleteOffLineActivity"`
  VendorSpecific2 byte `v8:"vendorSpecific2"`
  OffLineDataCollectionCapability byte `v8:"offLineDataCollectionCapability"`
  SMARTCapability [2]byte `v8:"sMARTCapability"`
  ErrorLoggingCapability byte `v8:"errorLoggingCapability"`
  VendorSpecific3 byte `v8:"vendorSpecific3"`
  ShortTestPollingInterval byte `v8:"shortTestPollingInterval"`
  ExtendedTestPollingInterval byte `v8:"extendedTestPollingInterval"`
  Reserved [12]byte `v8:"reserved"`
  VendorSpecific4 [125]byte `v8:"vendorSpecific4"`
  Checksum byte `v8:"checksum"`
}
