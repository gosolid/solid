//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

/*
interface QuartzCore.CATransformLayer : QuartzCore.CALayer
*/

type CATransformLayer struct {
  *CALayer

}

