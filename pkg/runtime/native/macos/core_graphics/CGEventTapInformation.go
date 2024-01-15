//js:package native/macos/core-graphics
package core_graphics

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreGraphics.__CGEventTapInformation
*/

type CGEventTapInformation struct {
  EventTapID uint `v8:"eventTapID"`
  TapPoint int `v8:"tapPoint"`
  Options int `v8:"options"`
  EventsOfInterest uint64 `v8:"eventsOfInterest"`
  TappingProcess int `v8:"tappingProcess"`
  ProcessBeingTapped int `v8:"processBeingTapped"`
  Enabled func(...any) (any, error) `v8:"enabled"`
  MinUsecLatency float32 `v8:"minUsecLatency"`
  AvgUsecLatency float32 `v8:"avgUsecLatency"`
  MaxUsecLatency float32 `v8:"maxUsecLatency"`
}
