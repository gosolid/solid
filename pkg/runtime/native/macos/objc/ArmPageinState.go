//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__arm_pagein_state
*/

type ArmPageinState struct {
  PageinError int `v8:"pageinError"`
}
