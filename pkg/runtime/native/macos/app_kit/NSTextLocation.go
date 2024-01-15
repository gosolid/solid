//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTextLocation
*/

type NSTextLocation struct {

}

func (r *NSTextLocation) Compare() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

