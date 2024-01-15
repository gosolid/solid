//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFontMetrics
*/

type ATSFontMetrics struct {
  Version uint `v8:"version"`
  Ascent float64 `v8:"ascent"`
  Descent float64 `v8:"descent"`
  Leading float64 `v8:"leading"`
  AvgAdvanceWidth float64 `v8:"avgAdvanceWidth"`
  MaxAdvanceWidth float64 `v8:"maxAdvanceWidth"`
  MinLeftSideBearing float64 `v8:"minLeftSideBearing"`
  MinRightSideBearing float64 `v8:"minRightSideBearing"`
  StemWidth float64 `v8:"stemWidth"`
  StemHeight float64 `v8:"stemHeight"`
  CapHeight float64 `v8:"capHeight"`
  XHeight float64 `v8:"xHeight"`
  ItalicAngle float64 `v8:"italicAngle"`
  UnderlinePosition float64 `v8:"underlinePosition"`
  UnderlineThickness float64 `v8:"underlineThickness"`
}
