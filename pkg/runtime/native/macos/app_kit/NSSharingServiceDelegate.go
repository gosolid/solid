//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSSharingServiceDelegate
*/

type NSSharingServiceDelegate struct {

}

func (r *NSSharingServiceDelegate) SharingService() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSharingServiceDelegate) AnchoringViewForSharingService() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

