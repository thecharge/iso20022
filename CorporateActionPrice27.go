package iso20022

// Specifies prices.
type CorporateActionPrice27 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat5Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice2Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *AmountPrice2 `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat6Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat7Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice27) AddCashInLieuOfSharePrice() *PriceFormat5Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat5Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice27) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice2Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice2Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice27) AddCashValueForTax() *AmountPrice2 {
	c.CashValueForTax = new(AmountPrice2)
	return c.CashValueForTax
}

func (c *CorporateActionPrice27) AddGenericCashPricePaidPerProduct() *PriceFormat6Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat6Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice27) AddGenericCashPriceReceivedPerProduct() *PriceFormat7Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat7Choice)
	return c.GenericCashPriceReceivedPerProduct
}
