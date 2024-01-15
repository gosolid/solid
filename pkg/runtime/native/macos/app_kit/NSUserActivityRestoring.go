//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSUserActivityRestoring
*/

type NSUserActivityRestoring struct {

}

func (r *NSUserActivityRestoring) RestoreUserActivityState() error {
  return fmt.Errorf("unimplemented")
}

