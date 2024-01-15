//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTextLayoutManagerDelegate
*/

type NSTextLayoutManagerDelegate struct {

}

func (r *NSTextLayoutManagerDelegate) TextLayoutManager() (*NSTextLayoutFragment, error) {
  return nil, fmt.Errorf("unimplemented")
}

