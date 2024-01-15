//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSDataDetector : Foundation.NSRegularExpression
*/

type NSDataDetector struct {
  *NSRegularExpression

}

func (r *NSDataDetector) DataDetectorWithTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDataDetector) InitWithTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDataDetector) CheckingTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

