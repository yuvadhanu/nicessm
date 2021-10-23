package services

import (
	"fmt"
	"nicessm-api-service/models"
)

//ACLAccess :
func (s *Service) ACLAccess(ctx *models.Context, userTypeID string) (*models.ACLAccess, error) {

	var aclAccess = new(models.ACLAccess)
	utma, err := s.GetSingleModuleUserType(ctx, userTypeID)
	if err != nil {
		return nil, err
	}
	// aclAccess.Module = utma

	if utma != nil {
		if len(utma.Modules) > 0 {
			for _, v := range utma.Modules {
				var utmenua *models.UserTypeMenuAccess
				aclAccess.ModuleAccess = append(aclAccess.ModuleAccess, v)

				utmenua, err := s.GetSingleUserTypeMenuAccess(ctx, userTypeID, v.UniqueID)
				if err == nil {
					// aclAccess.Menu = append(aclAccess.Menu, *utmenua)
					if len(utmenua.Module.Menus) > 0 {
						for _, v2 := range utmenua.Module.Menus {
							aclAccess.MenuAccess = append(aclAccess.MenuAccess, v2)

						}
					}
				}
				if err != nil {
					fmt.Println("err in menu find - menu - ", v.UniqueID, "user type")
				}

				uttaba, err := s.GetSingleUserTypeTabAccess(ctx, userTypeID, v.UniqueID)
				if err == nil {

					// aclAccess.Tab = append(aclAccess.Tab, *uttaba)

					if len(uttaba.Module.Tabs) > 0 {
						for _, v2 := range uttaba.Module.Tabs {
							aclAccess.TabAccess = append(aclAccess.TabAccess, v2)
						}
					}
				}
				if err != nil {
					fmt.Println("err in menu find - tab - ", v.UniqueID, "user type", userTypeID)
				}

				utfeaturea, err := s.GetSingleUserTypeFeatureAccess(ctx, userTypeID, v.UniqueID)
				if err == nil {
					// aclAccess.Feature = append(aclAccess.Feature, *utfeaturea)
					if len(utfeaturea.Module.Features) > 0 {
						for _, v2 := range utfeaturea.Module.Features {
							aclAccess.FeatureAccess = append(aclAccess.FeatureAccess, v2)
						}
					}
				}
				if err != nil {
					fmt.Println("err in menu find - feature - ", v.UniqueID, "user type", userTypeID)
				}
			}
		}
	}
	return aclAccess, nil
}
