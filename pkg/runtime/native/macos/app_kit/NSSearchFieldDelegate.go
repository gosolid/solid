//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSSearchFieldDelegate
*/

type NSSearchFieldDelegate struct {

}

func (r *NSSearchFieldDelegate) SearchFieldDidStartSearching() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldDelegate) SearchFieldDidEndSearching() error {
  return fmt.Errorf("unimplemented")
}

