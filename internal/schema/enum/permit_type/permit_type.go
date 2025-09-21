package permit_type

import "github.com/pkg/errors"

type PermitType string

func (p PermitType) String() string {
	return string(p)
}

const (
	Work         PermitType = "WORK"
	CompanyOwner PermitType = "COMPANY_OWNER"
	Freelance    PermitType = "FREELANCE"
	Student      PermitType = "STUDENT"
	Investor     PermitType = "INVESTOR"
	Family       PermitType = "FAMILY"
	Asylum       PermitType = "ASYLUM"
	Retirement   PermitType = "RETIREMENT"
	Other        PermitType = "OTHER"
	Unknown      PermitType = "UNKNOWN"
)

func New(str string) (PermitType, error) {
	switch str {
	case "WORK":
		return Work, nil
	case "COMPANY_OWNER":
		return CompanyOwner, nil
	case "FREELANCE":
		return Freelance, nil
	case "STUDENT":
		return Student, nil
	case "INVESTOR":
		return Investor, nil
	case "FAMILY":
		return Family, nil
	case "ASYLUM":
		return Asylum, nil
	case "RETIREMENT":
		return Retirement, nil
	case "OTHER":
		return Other, nil
	case "UNKNOWN":
		return Unknown, nil
	default:
		return "", errors.Errorf("can't define PermitType from: %v available values are: %v", str, Values())
	}
}

func Values() []PermitType {
	return []PermitType{
		Work,
		CompanyOwner,
		Freelance,
		Student,
		Investor,
		Family,
		Asylum,
		Retirement,
		Other,
		Unknown,
	}
}
