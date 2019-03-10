package repository

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"gomongo/src/modules/user/model"
)

// UserRepositoryMongo ...
type UserRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

// NewUserRepositoryMongo ...
func NewUserRepositoryMongo(db *mgo.Database, collection string) *UserRepositoryMongo {
	return &UserRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

// Save ...
func (r *UserRepositoryMongo) Save(user *model.User) error {
	err := r.db.C(r.collection).Insert(user)
	return err
}

// Update ...
func (r *UserRepositoryMongo) Update(id string, user *model.User) error {
	user.UpdatedAt = time.Now()
	err := r.db.C(r.collection).Update(bson.M{"id": id}, user)
	return err
}

// Delete ...
func (r *UserRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"id": id})
	return err
}

// FindByID ...
func (r *UserRepositoryMongo) FindByID(id string) (*model.User, error) {
	var user model.User
	err := r.db.C(r.collection).Find(bson.M{"id": id}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindAll ...
func (r *UserRepositoryMongo) FindAll() (model.Users, error) {
	var users model.Users

	err := r.db.C(r.collection).Find(bson.M{}).All(&users)

	if err != nil {
		return nil, err
	}
	return users, nil
}
