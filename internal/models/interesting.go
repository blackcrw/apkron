package models

type InterestingModel struct {
	Permissions 		  []InterestingPermissionAndroidModel
	Match                 []string
}

type InterestingPermissionAndroidModel struct {
	Name        string
	Description string
	Match       string
}