//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Foundation.NSAffineTransformStruct
*/

type NSAffineTransformStruct struct {
  M11 float64 `v8:"m11"`
  M12 float64 `v8:"m12"`
  M21 float64 `v8:"m21"`
  M22 float64 `v8:"m22"`
  TX float64 `v8:"tX"`
  TY float64 `v8:"tY"`
}
