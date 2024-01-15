//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ItlcRecord
*/

type ItlcRecord struct {
  ItlcSystem int16 `v8:"itlcSystem"`
  ItlcReserved int16 `v8:"itlcReserved"`
  ItlcFontForce byte `v8:"itlcFontForce"`
  ItlcIntlForce byte `v8:"itlcIntlForce"`
  ItlcOldKybd byte `v8:"itlcOldKybd"`
  ItlcFlags byte `v8:"itlcFlags"`
  ItlcIconOffset int16 `v8:"itlcIconOffset"`
  ItlcIconSide byte `v8:"itlcIconSide"`
  ItlcIconRsvd byte `v8:"itlcIconRsvd"`
  ItlcRegionCode int16 `v8:"itlcRegionCode"`
  ItlcSysFlags int16 `v8:"itlcSysFlags"`
  ItlcReserved4 [32]byte `v8:"itlcReserved4"`
}
