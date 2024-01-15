//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSKeyedArchiver : Foundation.NSCoder
*/

type NSKeyedArchiver struct {
  *NSCoder

}

func (r *NSKeyedArchiver) ArchivedDataWithRootObject() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) SetClassName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeInt32() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeFloat() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) SetRequiresSecureCoding() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) FinishEncoding() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) ClassNameForClass() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeConditionalObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeInt() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeInt64() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) InitRequiringSecureCoding() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) ArchiveRootObject() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeBool() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) OutputFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodedData() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) InitForWritingWithMutableData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeDouble() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeBytes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) SetOutputFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) RequiresSecureCoding() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

