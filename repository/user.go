package repository

import (
	"context"
	"rentify/entity"

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

func (self *UserRepo) GetByID(ctx context.Context, id string) error {
	var user entity.User
	err := self.coll.Find(ctx, bson.M{"id": id}).One(&user)
	if err != nil {
		return err
	}
	return nil
}
