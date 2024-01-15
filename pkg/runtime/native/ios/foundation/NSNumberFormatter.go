//js:package native/ios/foundation
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

func (r *NSNumberFormatter) SetMaximumIntegerDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) GetObjectValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) DefaultFormatterBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetTextAttributesForPositiveValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NegativeInfinitySymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPaddingPosition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetRoundingMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) Maximum() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) IsLenient() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) GeneratesDecimalNumbers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetNegativeFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) CurrencyDecimalSeparator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPositiveSuffix() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PercentSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PaddingCharacter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) UsesGroupingSeparator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetNegativeInfinitySymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMinimumIntegerDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMaximumSignificantDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) InternationalCurrencySymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetSecondaryGroupingSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMultiplier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) LocalizedStringFromNumber() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetDefaultFormatterBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetAllowsFloats() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) TextAttributesForNil() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NegativeSuffix() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) RoundingMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetLenient() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMinimumSignificantDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetTextAttributesForNegativeValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetUsesGroupingSeparator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) CurrencyCode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) UsesSignificantDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetZeroSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPerMillSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) MinimumFractionDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) TextAttributesForZero() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) IsPartialStringValidationEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetUsesSignificantDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetAlwaysShowsDecimalSeparator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPositiveInfinitySymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetCurrencyCode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetInternationalCurrencySymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPercentSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) StringFromNumber() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetTextAttributesForPositiveInfinity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PositivePrefix() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) RoundingIncrement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PositiveFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetNotANumberSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PositiveInfinitySymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PlusSign() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) MinimumIntegerDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) CurrencySymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPaddingCharacter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMinimumFractionDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) Locale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetFormatterBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NotANumberSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) TextAttributesForNotANumber() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) TextAttributesForPositiveInfinity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) AllowsFloats() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) MaximumIntegerDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) MaximumFractionDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetNilSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetFormattingContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) CurrencyGroupingSeparator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPositiveFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) GroupingSeparator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NumberStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetNegativeSuffix() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) MaximumSignificantDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) FormattingContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetTextAttributesForNegativeInfinity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMinimum() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) TextAttributesForPositiveValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NilSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetRoundingIncrement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) Minimum() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetFormatWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetGroupingSeparator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) TextAttributesForNegativeInfinity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) ExponentSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetExponentSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) Multiplier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SecondaryGroupingSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PaddingPosition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMaximumFractionDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetGeneratesDecimalNumbers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NegativeFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) TextAttributesForNegativeValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetTextAttributesForZero() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPlusSign() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMaximum() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetCurrencySymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetLocale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) FormatterBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetGroupingSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetCurrencyGroupingSeparator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetDecimalSeparator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NegativePrefix() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetMinusSign() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) MinimumSignificantDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetNegativePrefix() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PerMillSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) GroupingSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) FormatWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) NumberFromString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) DecimalSeparator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) ZeroSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) AlwaysShowsDecimalSeparator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetNumberStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetCurrencyDecimalSeparator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPositivePrefix() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetPartialStringValidationEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetTextAttributesForNil() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) MinusSign() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) SetTextAttributesForNotANumber() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNumberFormatter) PositiveSuffix() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

