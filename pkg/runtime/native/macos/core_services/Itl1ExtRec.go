//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.Itl1ExtRec
*/

type Itl1ExtRec struct {
  Base Intl1Rec `v8:"base"`
  Version int16 `v8:"version"`
  Format int16 `v8:"format"`
  CalendarCode int16 `v8:"calendarCode"`
  ExtraDaysTableOffset int `v8:"extraDaysTableOffset"`
  ExtraDaysTableLength int `v8:"extraDaysTableLength"`
  ExtraMonthsTableOffset int `v8:"extraMonthsTableOffset"`
  ExtraMonthsTableLength int `v8:"extraMonthsTableLength"`
  AbbrevDaysTableOffset int `v8:"abbrevDaysTableOffset"`
  AbbrevDaysTableLength int `v8:"abbrevDaysTableLength"`
  AbbrevMonthsTableOffset int `v8:"abbrevMonthsTableOffset"`
  AbbrevMonthsTableLength int `v8:"abbrevMonthsTableLength"`
  ExtraSepsTableOffset int `v8:"extraSepsTableOffset"`
  ExtraSepsTableLength int `v8:"extraSepsTableLength"`
  Tables [1]int16 `v8:"tables"`
}
