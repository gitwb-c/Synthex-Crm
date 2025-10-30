package rule

import (
	"context"
	"errors"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gitwb-c/crm.saas/backend/internal/viewer"
	"github.com/google/uuid"
)

func FilterTenantMutation() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
			op := mutation.Op()

			view := viewer.FromContext(ctx)
			if view.Signature != nil {
				switch *view.Signature {
				case viewer.Bootstrap:
					return next.Mutate(ctx, mutation)
				}
			}
			if view.TenantID == uuid.Nil {
				return nil, errors.New("missing tenant")
			}

			switch op {
			case ent.OpCreate:
				current, exists := mutation.Field("tenant_id")
				if exists {
					if current.(uuid.UUID) != view.TenantID {
						return nil, errors.New("mismatched tenant")
					}
				} else {
					if err := mutation.SetField("tenant_id", view.TenantID); err != nil {
						return nil, err
					}
				}
			case ent.OpUpdate, ent.OpUpdateOne, ent.OpDelete, ent.OpDeleteOne:
				mp, ok := mutation.(interface {
					WhereP(...func(*sql.Selector))
				})
				if !ok {
					return nil, errors.New("unsupported mutation type")
				}
				mp.WhereP(func(s *sql.Selector) {
					s.Where(sql.EQ(s.C("tenant_id"), view.TenantID))
				})
			}

			return next.Mutate(ctx, mutation)
		})
	}
}
