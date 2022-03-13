package text

import "github.com/blackcrw/apkron/internal/models"

func ContainsPermissionByName(model *[]models.InterestingPermissionAndroidModel, permission_name string) bool {
	for _, x := range *model {
		if x.Name == permission_name {
			return true
		}
	}
	
	return false
}