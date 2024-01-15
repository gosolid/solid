//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSScanner : objc.NSObject
*/

type NSScanner struct {
  *objc.NSObject

}

func (r *NSScanner) SetCharactersToBeSkipped() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScanner) CaseSensitive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScanner) SetCaseSensitive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScanner) SetLocale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScanner) ScanLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScanner) SetScanLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScanner) CharactersToBeSkipped() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScanner) InitWithString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScanner) String() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScanner) Locale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

