//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IODetailedTimingInformationV1
*/

type IODetailedTimingInformationV1 struct {
  PixelClock uint `v8:"pixelClock"`
  HorizontalActive uint `v8:"horizontalActive"`
  HorizontalBlanking uint `v8:"horizontalBlanking"`
  HorizontalBorder uint `v8:"horizontalBorder"`
  HorizontalSyncOffset uint `v8:"horizontalSyncOffset"`
  HorizontalSyncWidth uint `v8:"horizontalSyncWidth"`
  VerticalActive uint `v8:"verticalActive"`
  VerticalBlanking uint `v8:"verticalBlanking"`
  VerticalBorder uint `v8:"verticalBorder"`
  VerticalSyncOffset uint `v8:"verticalSyncOffset"`
  VerticalSyncWidth uint `v8:"verticalSyncWidth"`
}
