package rulehelpers

import (
	"entgo.io/ent/privacy"
	"github.com/gitwb-c/crm.saas/backend/internal/viewer"
)

func VerifyAccess(neededPermissions []string, view viewer.UserViewer) error {
	if len(neededPermissions) == 0 {
		return nil
	}
	if view.Permissions == nil || len(*view.Permissions) == 0 {
		return privacy.Denyf("viewer do not have any permissions")
	}
	if len(*view.Permissions) < len(neededPermissions) {
		return privacy.Denyf("viewer do not have sufficient permissions")
	}

	permissionMap := make(map[string]struct{})
	for _, perm := range *view.Permissions {
		permissionMap[perm] = struct{}{}
	}

	for _, neededPerm := range neededPermissions {
		if _, exists := permissionMap[neededPerm]; !exists {
			return privacy.Denyf("viewer do not have sufficient permissions")
		}
	}
	return nil
}
