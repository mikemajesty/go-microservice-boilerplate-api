package utils

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IEntity interface {
	GetID() any
	SetID(id any)
	SetCreatedAt()
	SetUpdatedAt()
}

type IEntityID interface {
	~*primitive.ObjectID | uint | string
}

type Entity[T IEntityID] struct {
	ID        T         `bson:"_id" json:"id" gorm:"primarykey"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func (b *Entity[T]) GetID() any {
	return b.ID
}

func (b *Entity[T]) SetID(id any) {
	b.ID = id.(T)
}

func (b *Entity[T]) SetCreatedAt() {
	b.CreatedAt = time.Now()
}

func (b *Entity[T]) SetUpdatedAt() {
	b.UpdatedAt = time.Now()
}
