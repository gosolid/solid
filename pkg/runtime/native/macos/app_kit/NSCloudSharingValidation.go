//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSCloudSharingValidation
*/

type NSCloudSharingValidation struct {

}

func (r *NSCloudSharingValidation) CloudShareForUserInterfaceItem() (*CKShare, error) {
  return nil, fmt.Errorf("unimplemented")
}

