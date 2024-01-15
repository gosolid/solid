//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOOutputQueueStats
*/

type IOOutputQueueStats struct {
  Capacity uint `v8:"capacity"`
  Size uint `v8:"size"`
  PeakSize uint `v8:"peakSize"`
  DropCount uint `v8:"dropCount"`
  OutputCount uint `v8:"outputCount"`
  RetryCount uint `v8:"retryCount"`
  StallCount uint `v8:"stallCount"`
  Reserved [4]uint `v8:"reserved"`
}
