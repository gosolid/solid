//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSSoundDelegate
*/

type NSSoundDelegate struct {

}

func (r *NSSoundDelegate) Sound() error {
  return fmt.Errorf("unimplemented")
}

