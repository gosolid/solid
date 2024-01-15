//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMassFormatter : Foundation.NSFormatter
*/

type NSMassFormatter struct {
  *NSFormatter

}

func (r *NSMassFormatter) GetObjectValue() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) NumberFormatter() (*NSNumberFormatter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) SetNumberFormatter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) UnitStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) SetUnitStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) SetForPersonMassUse() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) StringFromValue() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) StringFromKilograms() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) UnitStringFromValue() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) UnitStringFromKilograms() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) IsForPersonMassUse() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

