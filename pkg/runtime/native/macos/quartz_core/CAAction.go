//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol QuartzCore.CAAction
*/

type CAAction struct {

}

func (r *CAAction) RunActionForKey() error {
  return fmt.Errorf("unimplemented")
}

