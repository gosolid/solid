//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.InstrumentChunk
*/

type InstrumentChunk struct {
  CkID uint `v8:"ckID"`
  CkSize int `v8:"ckSize"`
  BaseFrequency byte `v8:"baseFrequency"`
  Detune byte `v8:"detune"`
  LowFrequency byte `v8:"lowFrequency"`
  HighFrequency byte `v8:"highFrequency"`
  LowVelocity byte `v8:"lowVelocity"`
  HighVelocity byte `v8:"highVelocity"`
  Gain int16 `v8:"gain"`
  SustainLoop AIFFLoop `v8:"sustainLoop"`
  ReleaseLoop AIFFLoop `v8:"releaseLoop"`
}
