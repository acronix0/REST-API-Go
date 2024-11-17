package repository

import (
	"context"
	"strconv"
	"time"

	"github.com/acronix0/REST-API-Go/internal/cache"
)

type UserRoleRepo struct{
	cache cache.CacheProvider
}
func NewUserRoleRepo(cache cache.CacheProvider) *UserRoleRepo{
	return &UserRoleRepo{cache: cache}
}

func (r *UserRoleRepo) Get(ctx context.Context, id int) (string, error){
	userRole, err := r.cache.Get(ctx,strconv.Itoa(id))
  if err!= nil {
    return "", err
  }
  return userRole, nil
}

func (r *UserRoleRepo) Set(ctx context.Context, id int, userRole string) error{
  return r.cache.Set(ctx, strconv.Itoa(id), userRole, time.Hour*3)
}

func (r *UserRoleRepo) Delete(ctx context.Context, id int) error{
  return r.cache.Delete(ctx, strconv.Itoa(id))
}