package rule

import (
	"context"

	"entgo.io/ent/privacy"
	rulehelpers "github.com/gitwb-c/crm.saas/backend/internal/rule/rule_helpers"
	"github.com/gitwb-c/crm.saas/backend/internal/viewer"
)

func MutationRules(permissions map[viewer.QueryType][]string) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		querytype, err := rulehelpers.PrivacyValidations(ctx, view)
		if err != nil {
			return err
		}
		neededPermissions := permissions[querytype]
		if err := rulehelpers.VerifyAccess(neededPermissions, view); err != nil {
			return err
		}
		return privacy.Allow
	})
}

func QueryRules(permissions map[viewer.QueryType][]string) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		querytype, err := rulehelpers.PrivacyValidations(ctx, view)
		if err != nil {
			return err
		}
		if querytype != viewer.Read {
			return privacy.Denyf("query type mismatch")
		}

		neededPermissions := permissions[querytype]
		if err := rulehelpers.VerifyAccess(neededPermissions, view); err != nil {
			return err
		}
		return privacy.Allow
	})
}
