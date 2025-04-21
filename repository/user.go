package repository

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/nara-ryoya/sqlboiler-mysql-view/models"
)

func ListUsers(ctx context.Context, exec boil.ContextExecutor) ([]*models.UsersView, error) {
	return models.UsersViews().All(ctx, exec)
}
