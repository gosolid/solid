//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSRegularExpression : objc.NSObject
*/

type NSRegularExpression struct {
  *objc.NSObject

}

func (r *NSRegularExpression) Pattern() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRegularExpression) Options() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRegularExpression) NumberOfCaptureGroups() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRegularExpression) RegularExpressionWithPattern() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRegularExpression) InitWithPattern() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRegularExpression) EscapedPatternForString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

