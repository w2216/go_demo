// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/add",
				Handler: addHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/edit",
				Handler: editHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/del",
				Handler: delHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/list",
				Handler: listHandler(serverCtx),
			},
		},
	)
}
