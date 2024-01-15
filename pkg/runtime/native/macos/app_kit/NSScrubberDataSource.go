//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSScrubberDataSource
*/

type NSScrubberDataSource struct {

}

func (r *NSScrubberDataSource) NumberOfItemsForScrubber() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrubberDataSource) Scrubber() (*NSScrubberItemView, error) {
  return nil, fmt.Errorf("unimplemented")
}

