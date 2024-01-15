//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTouchBarProvider
*/

type NSTouchBarProvider struct {

}

func (r *NSTouchBarProvider) TouchBar() (*NSTouchBar, error) {
  return nil, fmt.Errorf("unimplemented")
}

