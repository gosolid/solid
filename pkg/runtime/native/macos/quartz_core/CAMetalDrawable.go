//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol QuartzCore.CAMetalDrawable
*/

type CAMetalDrawable struct {

}

func (r *CAMetalDrawable) Texture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDrawable) Layer() (*CAMetalLayer, error) {
  return nil, fmt.Errorf("unimplemented")
}

