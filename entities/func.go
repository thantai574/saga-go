package entities

import "context"

type FuncSaga func(ctx context.Context) error
