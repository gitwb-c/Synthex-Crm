package rulehelpers

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent/privacy"
	"github.com/gitwb-c/crm.saas/backend/internal/graphql/graph/graph_helpers"
	"github.com/gitwb-c/crm.saas/backend/internal/viewer"
)

func PrivacyValidations(ctx context.Context, view viewer.UserViewer) (viewer.QueryType, error) {
	if view.Permissions == nil {
		return "", privacy.Denyf("missing viewer permisions")
	}
	valid, querytype := graph_helpers.GetQueryTypeCtx(ctx)
	if !valid {
		return "", privacy.Denyf("query type not verified")
	}
	if querytype == "" {
		return "", privacy.Denyf("missing viewer query type")
	}
	return viewer.QueryType(querytype), nil
}
