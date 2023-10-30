package repository

import (
	"context"
	"rentify/domain/entity"
	"rentify/exception"

	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

type UserRepo struct {
	coll *qmgo.Collection
}

func NewUserRepo(coll *qmgo.Collection) UserRepo {
	return UserRepo{
		coll: coll,
	}
}

func (slf *UserRepo) Create(ctx context.Context, user *entity.User) error {
	_, err := slf.coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (slf *UserRepo) Update(ctx context.Context, update *entity.User) error {
	err := slf.coll.UpdateOne(ctx, bson.M{"id": update.ID}, bson.M{"$set": update})
	if err != nil {
		return err
	}

	return nil
}

func (slf *UserRepo) Patch(ctx context.Context, id string, patch_payload bson.M) error {
	err := slf.coll.UpdateOne(
		ctx,
		bson.M{"id": id},
		bson.M{"$set": patch_payload},
	)
	if err != nil {
		return err
	}

	return nil
}

func (slf *UserRepo) GetByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	err := slf.coll.Find(ctx, bson.M{"id": id}).One(&user)
	if err == qmgo.ErrNoSuchDocuments {
		return nil, exception.DbObjNotFound
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (slf *UserRepo) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	err := slf.coll.Find(ctx, bson.M{"username": username}).One(&user)
	if err == qmgo.ErrNoSuchDocuments {
		return nil, exception.DbObjNotFound
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (slf *UserRepo) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := slf.coll.Find(ctx, bson.M{"email": email}).One(&user)
	if err == qmgo.ErrNoSuchDocuments {
		return nil, exception.DbObjNotFound
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (slf *UserRepo) GetList(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	err := slf.coll.Find(ctx, bson.M{}).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
