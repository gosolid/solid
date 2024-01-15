//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

/*
struct QuartzCore.CAFrameRateRange
*/

type CAFrameRateRange struct {
  Minimum float32 `v8:"minimum"`
  Maximum float32 `v8:"maximum"`
  Preferred float32 `v8:"preferred"`
}
