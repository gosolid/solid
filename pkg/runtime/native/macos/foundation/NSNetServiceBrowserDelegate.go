//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSNetServiceBrowserDelegate
*/

type NSNetServiceBrowserDelegate struct {

}

func (r *NSNetServiceBrowserDelegate) NetServiceBrowserWillSearch() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetServiceBrowserDelegate) NetServiceBrowserDidStopSearch() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetServiceBrowserDelegate) NetServiceBrowser() error {
  return fmt.Errorf("unimplemented")
}

