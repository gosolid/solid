//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ItlbExtRecord
*/

type ItlbExtRecord struct {
  Base ItlbRecord `v8:"base"`
  ItlbLocalSize int `v8:"itlbLocalSize"`
  ItlbMonoFond int16 `v8:"itlbMonoFond"`
  ItlbMonoSize int16 `v8:"itlbMonoSize"`
  ItlbPrefFond int16 `v8:"itlbPrefFond"`
  ItlbPrefSize int16 `v8:"itlbPrefSize"`
  ItlbSmallFond int16 `v8:"itlbSmallFond"`
  ItlbSmallSize int16 `v8:"itlbSmallSize"`
  ItlbSysFond int16 `v8:"itlbSysFond"`
  ItlbSysSize int16 `v8:"itlbSysSize"`
  ItlbAppFond int16 `v8:"itlbAppFond"`
  ItlbAppSize int16 `v8:"itlbAppSize"`
  ItlbHelpFond int16 `v8:"itlbHelpFond"`
  ItlbHelpSize int16 `v8:"itlbHelpSize"`
  ItlbValidStyles byte `v8:"itlbValidStyles"`
  ItlbAliasStyle byte `v8:"itlbAliasStyle"`
}
