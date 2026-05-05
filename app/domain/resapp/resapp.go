package resapp

import (
	"context"
	"net/http"

	"github.com/OJPARKINSON/gotel/app/sdk/errs"
	"github.com/OJPARKINSON/gotel/business/domain/resbus"
	"github.com/OJPARKINSON/gotel/foundation/web"
)

type app struct {
	responseBus resbus.ExtBusiness
}

func newApp(responseBus resbus.ExtBusiness) *app {
	return &app{
		responseBus: responseBus,
	}
}

func (a *app) create(ctx context.Context, r *http.Request) web.Encoder {
	var app NewReservation

	if err := web.Decode(r, &app); err != nil {
		return errs.New(errs.InvalidArgument, err)
	}

	nr, err := toBusNewReservation(app) // app -> bus conversion
	if err != nil {
		return errs.New(errs.InvalidArgument, err)
	}

	res, err := a.responseBus.Create(ctx, nr)
	if err != nil {
		return errs.Newf(errs.Internal, "create: %s", err)
	}

	return toAppReservation(res)
}
