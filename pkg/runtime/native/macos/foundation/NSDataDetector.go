//js:package native/macos/foundation
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

func (r *NSDataDetector) DataDetectorWithTypes() (*NSDataDetector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDataDetector) InitWithTypes() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDataDetector) CheckingTypes() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

