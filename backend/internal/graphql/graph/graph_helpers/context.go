package graph_helpers

import (
	"context"
)

type key int

const graphkey key = 1

func AddQueryTypeCtx(ctx context.Context, querytype string) context.Context {
	return context.WithValue(ctx, graphkey, querytype)
}

func GetQueryTypeCtx(ctx context.Context) (bool, string) {
	query := ctx.Value(graphkey)
	if query == nil {
		return false, ""

	}
	querytypeStr, ok := query.(string)
	if !ok {
		return false, ""
	}
	return true, querytypeStr
}
