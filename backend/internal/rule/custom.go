package rule

import (
	"github.com/gitwb-c/crm.saas/backend/internal/ent/privacy"
	"github.com/gitwb-c/crm.saas/backend/internal/viewer"
)

func MutationRules(permissions map[viewer.QueryType][]string) privacy.QueryMutationRule {
	return privacy.AlwaysDenyRule()
}

func QueryRules(permissions []string) privacy.QueryMutationRule {
	return privacy.AlwaysDenyRule()
}
