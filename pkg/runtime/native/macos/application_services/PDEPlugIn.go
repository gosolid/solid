//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol ApplicationServices.PDEPlugIn
*/

type PDEPlugIn struct {

}

func (r *PDEPlugIn) InitWithBundle() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *PDEPlugIn) PDEPanelsForType() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

