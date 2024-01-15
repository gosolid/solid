//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol QuartzCore.CAMetalDisplayLinkDelegate
*/

type CAMetalDisplayLinkDelegate struct {

}

func (r *CAMetalDisplayLinkDelegate) MetalDisplayLink() error {
  return fmt.Errorf("unimplemented")
}

