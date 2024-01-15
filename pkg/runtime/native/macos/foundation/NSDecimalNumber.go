//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSDecimalNumber : Foundation.NSNumber
*/

type NSDecimalNumber struct {
  *NSNumber

}

func (r *NSDecimalNumber) Compare() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) DoubleValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) DecimalNumberByAdding() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) InitWithMantissa() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) DecimalValue() (NSDecimal, error) {
  return NSDecimal{}, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) NotANumber() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) InitWithDecimal() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) DecimalNumberByRoundingAccordingToBehavior() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) SetDefaultBehavior() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) DecimalNumberByRaisingToPower() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) DecimalNumberWithMantissa() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) DecimalNumberWithDecimal() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) DecimalNumberByMultiplyingByPowerOf10() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) One() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) MaximumDecimalNumber() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) DescriptionWithLocale() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) DecimalNumberBySubtracting() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) DecimalNumberByDividingBy() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) DecimalNumberWithString() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) DecimalNumberByMultiplyingBy() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) Zero() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) MinimumDecimalNumber() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) DefaultBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) ObjCType() (byte, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumber) InitWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

