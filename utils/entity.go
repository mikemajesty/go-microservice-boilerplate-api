package utils_entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IEntity interface {
	GetID() primitive.ObjectID
	SetID(id primitive.ObjectID)
	SetCreatedAt()
	SetUpdatedAt()
}

type Entity struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (b *Entity) GetID() primitive.ObjectID {
	return b.ID
}

func (b *Entity) SetID(id primitive.ObjectID) {
	b.ID = id
}

func (b *Entity) SetCreatedAt() {
	b.CreatedAt = time.Now()
}

func (b *Entity) SetUpdatedAt() {
	b.UpdatedAt = time.Now()
}
