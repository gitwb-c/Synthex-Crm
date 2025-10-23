package hooks

import (
	"context"
	"errors"

	"entgo.io/ent"
)

func TenantFilterHook(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
		tenantID, ok := ctx.Value("tenant_id").(int)
		if !ok {
			return nil, errors.New("tenant_id not found in ctx")
		}

		return next.Query(ctx, tenantID)
	})
}
