//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.Intl1Rec
*/

type Intl1Rec struct {
  Days [7][16]byte `v8:"days"`
  Months [12][16]byte `v8:"months"`
  SuppressDay byte `v8:"suppressDay"`
  LngDateFmt byte `v8:"lngDateFmt"`
  DayLeading0 byte `v8:"dayLeading0"`
  AbbrLen byte `v8:"abbrLen"`
  St0 [4]byte `v8:"st0"`
  St1 [4]byte `v8:"st1"`
  St2 [4]byte `v8:"st2"`
  St3 [4]byte `v8:"st3"`
  St4 [4]byte `v8:"st4"`
  Intl1Vers int16 `v8:"intl1Vers"`
  LocalRtn [1]int16 `v8:"localRtn"`
}
