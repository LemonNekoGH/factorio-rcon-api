package factorioapi

import (
	factorioapiv1 "github.com/nekomeowww/factorio-rcon-api/apis/factorioapi/v1"
	consolev1 "github.com/nekomeowww/factorio-rcon-api/internal/grpc/services/factorioapi/v1/console"
	grpcpkg "github.com/nekomeowww/factorio-rcon-api/pkg/grpc"
	"go.uber.org/fx"
	"google.golang.org/grpc/reflection"
)

func Modules() fx.Option {
	return fx.Options(
		fx.Provide(NewFactorioAPI()),
		fx.Provide(consolev1.NewConsoleService()),
	)
}

type NewFactorioAPIParams struct {
	fx.In

	Console *consolev1.ConsoleService
}

type FactorioAPI struct {
	params *NewFactorioAPIParams
}

func NewFactorioAPI() func(params NewFactorioAPIParams) *FactorioAPI {
	return func(params NewFactorioAPIParams) *FactorioAPI {
		return &FactorioAPI{params: &params}
	}
}

func (c *FactorioAPI) Register(r *grpcpkg.Register) {
	r.RegisterHttpHandlers([]grpcpkg.HttpHandler{
		factorioapiv1.RegisterConsoleServiceHandler,
	})

	r.RegisterGrpcService(func(s reflection.GRPCServer) {
		factorioapiv1.RegisterConsoleServiceServer(s, c.params.Console)
	})
}
