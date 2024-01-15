//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSCloudSharingServiceDelegate
*/

type NSCloudSharingServiceDelegate struct {

}

func (r *NSCloudSharingServiceDelegate) SharingService() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCloudSharingServiceDelegate) OptionsForSharingService() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

