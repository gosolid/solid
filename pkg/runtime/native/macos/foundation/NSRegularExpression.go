//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSRegularExpression : objc.NSObject
*/

type NSRegularExpression struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSRegularExpression) RegularExpressionWithPattern() (*NSRegularExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRegularExpression) InitWithPattern() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRegularExpression) EscapedPatternForString() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRegularExpression) Pattern() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRegularExpression) Options() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRegularExpression) NumberOfCaptureGroups() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

