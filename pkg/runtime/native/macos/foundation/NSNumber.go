//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSNumber : Foundation.NSValue
*/

type NSNumber struct {
  *NSValue

}

func (r *NSNumber) InitWithLong() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) InitWithDouble() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) IsEqualToNumber() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSNumber) UnsignedIntValue() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumber) UnsignedIntegerValue() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumber) InitWithUnsignedInteger() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) ShortValue() (int16, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumber) UnsignedShortValue() (uint16, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumber) DescriptionWithLocale() (*objc.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) DoubleValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumber) IntegerValue() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumber) InitWithUnsignedInt() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) InitWithUnsignedLong() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) InitWithFloat() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) LongLongValue() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumber) UnsignedLongLongValue() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumber) BoolValue() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSNumber) InitWithUnsignedShort() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) InitWithBool() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) CharValue() (byte, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumber) IntValue() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumber) LongValue() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumber) InitWithChar() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) InitWithInt() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) InitWithUnsignedLongLong() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) UnsignedCharValue() (byte, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumber) FloatValue() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumber) StringValue() (*objc.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) InitWithUnsignedChar() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) InitWithShort() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) InitWithLongLong() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) InitWithInteger() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) Compare() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumber) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumber) UnsignedLongValue() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

