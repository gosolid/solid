//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOFBHDRMetaDataV1
*/

type IOFBHDRMetaDataV1 struct {
  DisplayPrimaryX0 uint16 `v8:"displayPrimaryX0"`
  DisplayPrimaryY0 uint16 `v8:"displayPrimaryY0"`
  DisplayPrimaryX1 uint16 `v8:"displayPrimaryX1"`
  DisplayPrimaryY1 uint16 `v8:"displayPrimaryY1"`
  DisplayPrimaryX2 uint16 `v8:"displayPrimaryX2"`
  DisplayPrimaryY2 uint16 `v8:"displayPrimaryY2"`
  DisplayPrimaryX uint16 `v8:"displayPrimaryX"`
  DisplayPrimaryY uint16 `v8:"displayPrimaryY"`
  DesiredLuminanceMax uint16 `v8:"desiredLuminanceMax"`
  DesiredLuminanceMin uint16 `v8:"desiredLuminanceMin"`
  DesiredLightLevelAvg uint16 `v8:"desiredLightLevelAvg"`
  DesiredLightLevelMax uint16 `v8:"desiredLightLevelMax"`
  ReservedA [5]uint64 `v8:"reservedA"`
}
