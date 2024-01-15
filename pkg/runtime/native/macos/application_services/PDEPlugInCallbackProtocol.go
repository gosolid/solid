//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
protocol ApplicationServices.PDEPlugInCallbackProtocol
*/

type PDEPlugInCallbackProtocol struct {

}

func (r *PDEPlugInCallbackProtocol) PageFormat() (*OpaquePMPageFormat, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *PDEPlugInCallbackProtocol) PMPrinter() (*OpaquePMPrinter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *PDEPlugInCallbackProtocol) PpdFile() (objc.PpdFileT, error) {
  return objc.PpdFileT{}, fmt.Errorf("unimplemented")
}

func (r *PDEPlugInCallbackProtocol) WillChangePPDOptionKeyValue() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *PDEPlugInCallbackProtocol) PrintSession() (*OpaquePMPrintSession, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *PDEPlugInCallbackProtocol) PrintSettings() (*OpaquePMPrintSettings, error) {
  return nil, fmt.Errorf("unimplemented")
}

