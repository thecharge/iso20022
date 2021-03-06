package iso20022

// Choice between an amount or a rate or an unspecified rate.
type InterestRateUsedForPaymentFormat1Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for interest rate used for payment.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus3 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateType12Code `xml:"NotSpcfdRate"`
}

func (i *InterestRateUsedForPaymentFormat1Choice) SetRate(value string) {
	i.Rate = (*PercentageRate)(&value)
}

func (i *InterestRateUsedForPaymentFormat1Choice) SetAmount(value, currency string) {
	i.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (i *InterestRateUsedForPaymentFormat1Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus3 {
	i.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus3)
	return i.RateTypeAndAmountAndRateStatus
}

func (i *InterestRateUsedForPaymentFormat1Choice) SetNotSpecifiedRate(value string) {
	i.NotSpecifiedRate = (*RateType12Code)(&value)
}
