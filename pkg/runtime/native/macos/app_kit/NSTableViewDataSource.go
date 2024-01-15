//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTableViewDataSource
*/

type NSTableViewDataSource struct {

}

func (r *NSTableViewDataSource) NumberOfRowsInTableView() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableViewDataSource) TableView() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

