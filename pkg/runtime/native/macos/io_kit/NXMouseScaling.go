//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.NXMouseScaling
*/

type NXMouseScaling struct {
  NumScaleLevels int `v8:"numScaleLevels"`
  ScaleThresholds [20]int16 `v8:"scaleThresholds"`
  ScaleFactors [20]int16 `v8:"scaleFactors"`
}
