//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTextStorageDelegate
*/

type NSTextStorageDelegate struct {

}

func (r *NSTextStorageDelegate) TextStorage() error {
  return fmt.Errorf("unimplemented")
}

