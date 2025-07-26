package utils

type Strings string

const (
	BackOfficeUser           = "system"
	AgentUser                = "Y"
	BusinessPartnerUser rune = '#'
	CallAndTrade             = "MFCNT"
	DummyFolio               = "_D"
	YES                      = "Y"
	NO                       = "N"
	DematHolding             = "D"
)

func (s Strings) IsCallAndTrade() bool {
	return s == CallAndTrade
}

func (s Strings) IsBackOfficeUser() bool {
	return s == BackOfficeUser
}

func (s Strings) IsAgent() bool {
	return s == AgentUser
}

func (s Strings) IsBusinessPartner() bool {
	var userIdPrefix = rune(s[0])
	return userIdPrefix == BusinessPartnerUser
}

func (s Strings) IsDummyFolio() bool {
	return s == DummyFolio
}

func (s Strings) IsYes() bool {
	return s == YES
}

func (s Strings) IsNo() bool {
	return s == NO
}

func (s Strings) IsDemat() bool {
	return s == DematHolding
}

func (s Strings) String() string {
	return string(s)
}

func (s Strings) IsNriUser() string {
	nriSubStr := s[:2]
	if nriSubStr == "65" || nriSubStr == "75" {
		return "Y"
	} else {
		return "N"
	}
}
