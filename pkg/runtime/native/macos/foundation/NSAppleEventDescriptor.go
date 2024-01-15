//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_services"
)

/*
interface Foundation.NSAppleEventDescriptor : objc.NSObject
*/

type NSAppleEventDescriptor struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSAppleEventDescriptor) DescriptorForKeyword() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) StringValue() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) ListDescriptor() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DescriptorAtIndex() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) InitWithEventClass() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) BooleanValue() (byte, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) Int32Value() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) EventClass() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DescriptorWithEnumCode() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DescriptorWithFileURL() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) SetDescriptor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) Data() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) SetParamDescriptor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) AttributeDescriptorForKeyword() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DescriptorType() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DoubleValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DescriptorWithDescriptorType() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DescriptorWithDate() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DescriptorWithProcessIdentifier() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) RemoveParamDescriptorWithKeyword() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) InsertDescriptor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) KeywordForDescriptorAtIndex() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) EventID() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DescriptorWithString() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) RecordDescriptor() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DescriptorWithBundleIdentifier() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) CoerceToDescriptorType() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) IsRecordDescriptor() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) NullDescriptor() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DescriptorWithInt32() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) SetAttributeDescriptor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) RemoveDescriptorWithKeyword() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) InitListDescriptor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) InitRecordDescriptor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) InitWithDescriptorType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) AeDesc() (core_services.AEDesc, error) {
  return core_services.AEDesc{}, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DescriptorWithBoolean() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DescriptorWithDouble() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DescriptorWithTypeCode() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) FileURLValue() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) RemoveDescriptorAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DateValue() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) TransactionID() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) CurrentProcessDescriptor() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) InitWithAEDescNoCopy() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) TypeCodeValue() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) NumberOfItems() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) ParamDescriptorForKeyword() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) SendEventWithOptions() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) EnumCodeValue() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) AppleEventWithEventClass() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) DescriptorWithApplicationURL() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleEventDescriptor) ReturnID() (int16, error) {
  return 0, fmt.Errorf("unimplemented")
}

