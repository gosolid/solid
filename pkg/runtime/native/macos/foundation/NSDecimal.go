//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Foundation.NSDecimal
*/

type NSDecimal struct {
  Exponent int `v8:"exponent"`
  Length uint `v8:"length"`
  IsNegative uint `v8:"isNegative"`
  IsCompact uint `v8:"isCompact"`
  Reserved uint `v8:"reserved"`
  Mantissa [8]uint16 `v8:"mantissa"`
}
