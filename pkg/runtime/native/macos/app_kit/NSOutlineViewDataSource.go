//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSOutlineViewDataSource
*/

type NSOutlineViewDataSource struct {

}

func (r *NSOutlineViewDataSource) OutlineView() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

