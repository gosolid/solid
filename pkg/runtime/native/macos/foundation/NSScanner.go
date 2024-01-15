//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSScanner : objc.NSObject
*/

type NSScanner struct {
  *objc.NSObject
  *NSCopying
}

func (r *NSScanner) CharactersToBeSkipped() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScanner) CaseSensitive() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScanner) SetCaseSensitive() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScanner) Locale() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScanner) InitWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScanner) String() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScanner) ScanLocation() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScanner) SetScanLocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScanner) SetLocale() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScanner) SetCharactersToBeSkipped() error {
  return fmt.Errorf("unimplemented")
}

