package models

import (
	"errors"
	"nicessm-api-service/constants"
)

func (u *User) ValidateAccessPrivilege() (bool, error) {
	switch u.AccessPrivilege.AccessLevel {
	case constants.ACCESSPRIVILEGESTATE:
		if len(u.AccessPrivilege.Districts) > 0 {
			return false, errors.New("invaild state access-districts not allowed")
		}
		if len(u.AccessPrivilege.States) > 0 {
			return false, errors.New("atleast one states required")
		}
		return true, nil
	case constants.ACCESSPRIVILEGEDISTRICT:
		if len(u.AccessPrivilege.States) == 0 || len(u.AccessPrivilege.States) > 1 {
			return false, errors.New("only one states allowed")
		}
		if len(u.AccessPrivilege.Districts) == 0 {
			return false, errors.New("atleast one district required")
		}
		return true, nil
	default:
		return false, errors.New("please mention access level")
	}
}
func (u *User) ValidateCallcenterAgent() (bool, error) {
	_, err := u.ValidateAccessPrivilege()
	if err != nil {
		return false, err
	}
	return true, nil
}
func (u *User) ValidateContentCreator() (bool, error) {
	_, err := u.ValidateAccessPrivilege()
	if err != nil {
		return false, err
	}
	if u.EducationalQualification == "" {
		return false, errors.New("please mention educational qualification")
	}
	if len(u.SubDomains) == 0 {
		return false, errors.New("atleast one subdomains required")
	}

	return true, nil
}
func (u *User) ValidateContentManager() (bool, error) {
	_, err := u.ValidateAccessPrivilege()
	if err != nil {
		return false, err
	}
	if u.EducationalQualification == "" {
		return false, errors.New("please mention educational qualification")
	}

	return true, nil
}
func (u *User) ValidateContentProvider() (bool, error) {
	_, err := u.ValidateAccessPrivilege()
	if err != nil {
		return false, err
	}
	if u.Organisation == "" {
		return false, errors.New("please mention organisation")
	}
	if u.Designation == "" {
		return false, errors.New("please mention designation")
	}
	if u.Officeaddress == "" {
		return false, errors.New("please mention officeaddress")
	}
	if u.Officenumber == 0 {
		return false, errors.New("please mention officenumber")
	}
	if u.OrganisationNatureofBusiness == "" {
		return false, errors.New("please mention OrganisationNatureofBusiness")
	}
	if u.KdExpertise == "" {
		return false, errors.New("please mention KdExpertise")
	}
	return true, nil
}
func (u *User) ValidateContentDisseminator() (bool, error) {
	_, err := u.ValidateAccessPrivilege()
	if err != nil {
		return false, err
	}
	if u.EducationalQualification == "" {
		return false, errors.New("please mention educational qualification")
	}

	return true, nil
}
func (u *User) ValidateFieldAgent() (bool, error) {
	_, err := u.ValidateAccessPrivilege()
	if err != nil {
		return false, err
	}
	if u.EducationalQualification == "" {
		return false, errors.New("please mention educational qualification")
	}
	if u.Organisation == "" {
		return false, errors.New("please mention organisation qualification")
	}
	if u.Designation == "" {
		return false, errors.New("please mention designation")
	}
	if u.Officeaddress == "" {
		return false, errors.New("please mention officeaddress")
	}
	if u.Officenumber == 0 {
		return false, errors.New("please mention officenumber")
	}
	if u.OrganisationNatureofBusiness == "" {
		return false, errors.New("please mention OrganisationNatureofBusiness")
	}

	return true, nil
}
func (u *User) ValidateLanguageTranslator() (bool, error) {
	_, err := u.ValidateAccessPrivilege()
	if err != nil {
		return false, err
	}
	if u.EducationalQualification == "" {
		return false, errors.New("please mention educational qualification")
	}
	if len(u.KnowledgeDomains) == 0 {
		return false, errors.New("please mention KnowledgeDomains")
	}

	if u.LanguageExpertise.IsZero() {
		return false, errors.New("please mention LanguageExpertise")
	}
	return true, nil
}
func (u *User) ValidateLanguageApprover() (bool, error) {
	_, err := u.ValidateAccessPrivilege()
	if err != nil {
		return false, err
	}
	if u.EducationalQualification == "" {
		return false, errors.New("please mention educational qualification")
	}
	if len(u.KnowledgeDomains) == 0 {
		return false, errors.New("please mention KnowledgeDomains")
	}
	if u.LanguageExpertise.IsZero() {
		return false, errors.New("please mention LanguageExpertise")
	}
	return true, nil
}
func (u *User) ValidateManagement() (bool, error) {
	_, err := u.ValidateAccessPrivilege()
	if err != nil {
		return false, err
	}
	if u.Organisation == "" {
		return false, errors.New("please mention organisation")
	}
	if u.Designation == "" {
		return false, errors.New("please mention designation")
	}
	return true, nil
}
func (u *User) ValidateModerator() (bool, error) {
	_, err := u.ValidateAccessPrivilege()
	if err != nil {
		return false, err
	}
	if u.EducationalQualification == "" {
		return false, errors.New("please mention educational qualification")
	}
	if len(u.KnowledgeDomains) == 0 {
		return false, errors.New("please mention KnowledgeDomains")
	}
	if len(u.SubDomains) == 0 {
		return false, errors.New("atleast one subdomains required")
	}
	return true, nil
}
func (u *User) ValidateSubjectMatterExpert() (bool, error) {
	_, err := u.ValidateAccessPrivilege()
	if err != nil {
		return false, err
	}
	if u.EducationalQualification == "" {
		return false, errors.New("please mention educational qualification")
	}
	if u.SubjectExpertise == "" {
		return false, errors.New("please mention  SubjectExpertise")
	}
	if u.Occupation == "" {
		return false, errors.New("please mention Occupation")
	}
	if len(u.KnowledgeDomains) == 0 {
		return false, errors.New("please mention KnowledgeDomains ")
	}
	if u.LanguageExpertise.IsZero() {
		return false, errors.New("please mention LanguageExpertise")
	}
	return true, nil
}
func (u *User) ValidateSystemAdmin() (bool, error) {
	_, err := u.ValidateAccessPrivilege()
	if err != nil {
		return false, err
	}
	if u.EducationalQualification == "" {
		return false, errors.New("please mention educational qualification")
	}
	return true, nil
}
func (u *User) ValidateVisitorAdmin() (bool, error) {
	_, err := u.ValidateAccessPrivilege()
	if err != nil {
		return false, err
	}
	if u.Organisation == "" {
		return false, errors.New("please mention organisation")
	}
	return true, nil
}
func (u *User) ValidateTrainer() (bool, error) {
	_, err := u.ValidateAccessPrivilege()
	if err != nil {
		return false, err
	}
	if u.EducationalQualification == "" {
		return false, errors.New("please mention educational qualification")
	}
	if u.Designation == "" {
		return false, errors.New("please mention Designation")
	}
	return true, nil
}
func (u *User) ValidateFieldAgentLead() (bool, error) {

	if u.EducationalQualification == "" {
		return false, errors.New("please mention educational qualification")
	}
	if u.Organisation == "" {
		return false, errors.New("please mention organisation qualification")
	}
	if u.Designation == "" {
		return false, errors.New("please mention designation")
	}
	if u.Officeaddress == "" {
		return false, errors.New("please mention officeaddress")
	}
	if u.Officenumber == 0 {
		return false, errors.New("please mention officenumber")
	}
	if u.OrganisationNatureofBusiness == "" {
		return false, errors.New("please mention OrganisationNatureofBusiness")
	}

	return true, nil
}
