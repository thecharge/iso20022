package iso20022

// Choice between a standard code or proprietary code to specify the type of intermediate securities distribution.
type IntermediateSecuritiesDistributionTypeFormat5Choice struct {

	// Standard code to specify the type of intermediate security distribution.
	Code *IntermediateSecurityDistributionType4Code `xml:"Cd"`

	// Proprietary identification of the type of intermediate security distribution.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (i *IntermediateSecuritiesDistributionTypeFormat5Choice) SetCode(value string) {
	i.Code = (*IntermediateSecurityDistributionType4Code)(&value)
}

func (i *IntermediateSecuritiesDistributionTypeFormat5Choice) AddProprietary() *GenericIdentification20 {
	i.Proprietary = new(GenericIdentification20)
	return i.Proprietary
}
