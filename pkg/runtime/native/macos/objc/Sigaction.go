//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__sigaction
*/

type Sigaction struct {
  SigactionU void `v8:"sigactionU"`
  SaTramp func(...any) (any, error) `v8:"saTramp"`
  SaMask uint `v8:"saMask"`
  SaFlags int `v8:"saFlags"`
}
