//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSNetServiceBrowser : objc.NSObject
*/

type NSNetServiceBrowser struct {
  *objc.NSObject

}

func (r *NSNetServiceBrowser) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetServiceBrowser) ScheduleInRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetServiceBrowser) RemoveFromRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetServiceBrowser) SearchForBrowsableDomains() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetServiceBrowser) SearchForRegistrationDomains() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetServiceBrowser) SearchForServicesOfType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetServiceBrowser) Stop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetServiceBrowser) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetServiceBrowser) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetServiceBrowser) IncludesPeerToPeer() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSNetServiceBrowser) SetIncludesPeerToPeer() error {
  return fmt.Errorf("unimplemented")
}

