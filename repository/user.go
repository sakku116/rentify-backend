package repository

import (
	"context"
	"rentify/entity"
	"rentify/exception"
	"rentify/helper"

	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

type UserRepo struct {
	coll *qmgo.Collection
}

func NewUserRepo(coll *qmgo.Collection) *UserRepo {
	return &UserRepo{
		coll: coll,
	}
}

func (self *UserRepo) Create(ctx context.Context, user *entity.User) error {
	_, err := self.coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (self *UserRepo) Update(ctx context.Context, user *entity.User) error {
	err := self.coll.UpdateOne(ctx, bson.M{"id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return err
	}

	return nil
}

func (self *UserRepo) Patch(ctx context.Context, id string, patch_payload *entity.User) error {
	err := self.coll.UpdateOne(
		ctx,
		bson.M{"id": id},
		helper.ParseBsonPatchStruct(patch_payload),
	)
	if err != nil {
		return err
	}

	return nil
}

func (self *UserRepo) GetByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	err := self.coll.Find(ctx, bson.M{"id": id}).One(&user)
	if err == qmgo.ErrNoSuchDocuments {
		return nil, exception.DbObjNotFound
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (self *UserRepo) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	err := self.coll.Find(ctx, bson.M{"username": username}).One(&user)
	if err == qmgo.ErrNoSuchDocuments {
		return nil, exception.DbObjNotFound
	} else if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (self *UserRepo) GetList(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	err := self.coll.Find(ctx, bson.M{}).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
