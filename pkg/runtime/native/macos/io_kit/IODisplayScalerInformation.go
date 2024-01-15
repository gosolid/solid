//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IODisplayScalerInformation
*/

type IODisplayScalerInformation struct {
  ReservedA [1]uint `v8:"reservedA"`
  Version uint `v8:"version"`
  ReservedB [2]uint `v8:"reservedB"`
  ScalerFeatures uint `v8:"scalerFeatures"`
  MaxHorizontalPixels uint `v8:"maxHorizontalPixels"`
  MaxVerticalPixels uint `v8:"maxVerticalPixels"`
  ReservedC [5]uint `v8:"reservedC"`
}
