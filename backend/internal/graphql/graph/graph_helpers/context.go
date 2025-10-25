package graph_helpers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type key int

const graphkey key = 1

func AddQueryTypeCtx(ctx context.Context, querytype string) context.Context {
	return context.WithValue(ctx, graphkey, querytype)
}

func GetQueryTypeCtx(ctx context.Context) (int, string, map[string]any) {
	query := ctx.Value(graphkey)
	if query == nil {
		return http.StatusInternalServerError, "", gin.H{"error": "query type not found"}

	}
	querytypeStr, ok := query.(string)
	if !ok {
		return http.StatusInternalServerError, "", gin.H{"error": "invalid query type"}
	}
	return http.StatusOK, querytypeStr, nil
}
