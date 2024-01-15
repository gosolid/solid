//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

/*
struct QuartzCore.CATransform3D
*/

type CATransform3D struct {
  M11 float64 `v8:"m11"`
  M12 float64 `v8:"m12"`
  M13 float64 `v8:"m13"`
  M14 float64 `v8:"m14"`
  M21 float64 `v8:"m21"`
  M22 float64 `v8:"m22"`
  M23 float64 `v8:"m23"`
  M24 float64 `v8:"m24"`
  M31 float64 `v8:"m31"`
  M32 float64 `v8:"m32"`
  M33 float64 `v8:"m33"`
  M34 float64 `v8:"m34"`
  M41 float64 `v8:"m41"`
  M42 float64 `v8:"m42"`
  M43 float64 `v8:"m43"`
  M44 float64 `v8:"m44"`
}
