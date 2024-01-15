//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSSplitViewDelegate
*/

type NSSplitViewDelegate struct {

}

func (r *NSSplitViewDelegate) SplitViewWillResizeSubviews() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewDelegate) SplitViewDidResizeSubviews() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewDelegate) SplitView() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

