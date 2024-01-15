//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.Intl0Rec
*/

type Intl0Rec struct {
  DecimalPt byte `v8:"decimalPt"`
  ThousSep byte `v8:"thousSep"`
  ListSep byte `v8:"listSep"`
  CurrSym1 byte `v8:"currSym1"`
  CurrSym2 byte `v8:"currSym2"`
  CurrSym3 byte `v8:"currSym3"`
  CurrFmt byte `v8:"currFmt"`
  DateOrder byte `v8:"dateOrder"`
  ShrtDateFmt byte `v8:"shrtDateFmt"`
  DateSep byte `v8:"dateSep"`
  TimeCycle byte `v8:"timeCycle"`
  TimeFmt byte `v8:"timeFmt"`
  MornStr [4]byte `v8:"mornStr"`
  EveStr [4]byte `v8:"eveStr"`
  TimeSep byte `v8:"timeSep"`
  Time1Suff byte `v8:"time1Suff"`
  Time2Suff byte `v8:"time2Suff"`
  Time3Suff byte `v8:"time3Suff"`
  Time4Suff byte `v8:"time4Suff"`
  Time5Suff byte `v8:"time5Suff"`
  Time6Suff byte `v8:"time6Suff"`
  Time7Suff byte `v8:"time7Suff"`
  Time8Suff byte `v8:"time8Suff"`
  MetricSys byte `v8:"metricSys"`
  Intl0Vers int16 `v8:"intl0Vers"`
}
