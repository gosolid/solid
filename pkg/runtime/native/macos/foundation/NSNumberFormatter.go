//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSNumberFormatter : Foundation.NSFormatter
*/

type NSNumberFormatter struct {
  *NSFormatter

}

func (r *NSNumberFormatter) ZeroSymbol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetZeroSymbol() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetTextAttributesForPositiveInfinity() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMaximumFractionDigits() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) IsPartialStringValidationEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) DecimalSeparator() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetRoundingMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) MinimumFractionDigits() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetCurrencyGroupingSeparator() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMultiplier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetTextAttributesForNotANumber() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetTextAttributesForNegativeInfinity() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMaximumIntegerDigits() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) TextAttributesForNegativeValues() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMinusSign() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) MaximumSignificantDigits() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPercentSymbol() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetGroupingSeparator() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMaximumSignificantDigits() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NegativeFormat() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) AlwaysShowsDecimalSeparator() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPaddingPosition() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMinimum() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) AllowsFloats() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetNegativePrefix() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPlusSign() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) RoundingIncrement() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) TextAttributesForNotANumber() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NegativeInfinitySymbol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) CurrencyCode() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SecondaryGroupingSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetSecondaryGroupingSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PositiveFormat() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) InternationalCurrencySymbol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PlusSign() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetFormatterBehavior() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NilSymbol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PositiveSuffix() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) MinimumSignificantDigits() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetTextAttributesForPositiveValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) CurrencyDecimalSeparator() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) MaximumIntegerDigits() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) IsLenient() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPartialStringValidationEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) DefaultFormatterBehavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) UsesGroupingSeparator() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetNegativeInfinitySymbol() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMinimumIntegerDigits() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetDecimalSeparator() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetAllowsFloats() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) TextAttributesForZero() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetNotANumberSymbol() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NegativePrefix() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetFormattingContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PaddingPosition() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) MaximumFractionDigits() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) Maximum() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) CurrencyGroupingSeparator() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) GroupingSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) StringFromNumber() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetTextAttributesForNil() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetFormatWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) GetObjectValue() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) Locale() (*NSLocale, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NegativeSuffix() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) MinusSign() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NumberFromString() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetInternationalCurrencySymbol() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) Multiplier() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) TextAttributesForPositiveInfinity() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PercentSymbol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) TextAttributesForNegativeInfinity() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) TextAttributesForPositiveValues() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NotANumberSymbol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetCurrencyCode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) CurrencySymbol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetLenient() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetUsesSignificantDigits() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMinimumSignificantDigits() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) LocalizedStringFromNumber() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMaximum() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) GroupingSeparator() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetRoundingIncrement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) ExponentSymbol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetUsesGroupingSeparator() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMinimumFractionDigits() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) FormatterBehavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PositivePrefix() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) FormatWidth() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) UsesSignificantDigits() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetTextAttributesForNegativeValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetNegativeFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetGeneratesDecimalNumbers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPositiveInfinitySymbol() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPositiveFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetNumberStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetCurrencyDecimalSeparator() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetTextAttributesForZero() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PositiveInfinitySymbol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPositiveSuffix() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetCurrencySymbol() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NumberStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetNegativeSuffix() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPerMillSymbol() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetExponentSymbol() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetDefaultFormatterBehavior() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetNilSymbol() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PerMillSymbol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetGroupingSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) Minimum() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetAlwaysShowsDecimalSeparator() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) MinimumIntegerDigits() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) TextAttributesForNil() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPositivePrefix() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PaddingCharacter() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) RoundingMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) FormattingContext() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) GeneratesDecimalNumbers() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPaddingCharacter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetLocale() error {
  return fmt.Errorf("unimplemented")
}

