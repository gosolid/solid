//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSNetServiceDelegate
*/

type NSNetServiceDelegate struct {

}

func (r *NSNetServiceDelegate) NetServiceWillResolve() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetServiceDelegate) NetServiceDidResolveAddress() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetServiceDelegate) NetServiceDidStop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetServiceDelegate) NetServiceWillPublish() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetServiceDelegate) NetServiceDidPublish() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetServiceDelegate) NetService() error {
  return fmt.Errorf("unimplemented")
}

