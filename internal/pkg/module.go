package pkg

import (
	"github.com/jacexh/go-miaobi/internal/pkg/database"
	"github.com/jacexh/go-miaobi/internal/pkg/grpcsrv"
	"github.com/jacexh/go-miaobi/internal/pkg/httpsrv"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"internal.pkg",
	fx.Provide(httpsrv.NewHTTPServer),
	fx.Provide(database.NewMySQLDriver),
	fx.Provide(grpcsrv.NewGRPCServ),
)
