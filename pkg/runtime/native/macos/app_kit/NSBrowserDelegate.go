//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSBrowserDelegate
*/

type NSBrowserDelegate struct {

}

func (r *NSBrowserDelegate) Browser() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowserDelegate) RootItemForBrowser() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowserDelegate) BrowserWillScroll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowserDelegate) BrowserDidScroll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowserDelegate) BrowserColumnConfigurationDidChange() error {
  return fmt.Errorf("unimplemented")
}

