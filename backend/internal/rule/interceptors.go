package rule

import (
	"context"
	"errors"
	"log"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gitwb-c/crm.saas/backend/internal/viewer"
	"github.com/google/uuid"
)

func FilterTenantQuery() ent.Interceptor {
	return ent.InterceptFunc(func(next ent.Querier) ent.Querier {
		return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
			view := viewer.FromContext(ctx)
			if view.TenantID == uuid.Nil {
				return nil, errors.New("missing tenant")
			}
			qp, ok := query.(interface {
				WhereP(...func(*sql.Selector))
			})
			if ok {
				qp.WhereP(func(s *sql.Selector) {
					s.Where(sql.EQ(s.C("tenant_id"), view.TenantID))
				})
			}

			return next.Query(ctx, query)
		})
	})
}

func FilterTenantMutation() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
			op := mutation.Op()

			view := viewer.FromContext(ctx)
			if view.TenantID == uuid.Nil {
				return nil, errors.New("missing tenant")
			}
			log.Print(view.TenantID)

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
			case ent.OpUpdate, ent.OpUpdateOne, ent.OpDelete:
				mp, ok := mutation.(interface {
					WhereP(...func(*sql.Selector))
				})
				if ok {
					mp.WhereP(func(s *sql.Selector) {
						s.Where(sql.EQ(s.C("tenant_id"), view.TenantID))
					})
				}
			}

			return next.Mutate(ctx, mutation)
		})
	}
}
