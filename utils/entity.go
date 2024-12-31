package utils

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EntityAdapter interface {
	GetID() any
	SetID(id any)
	SetCreatedAt()
	SetUpdatedAt()
}

type EntityIDAdapter interface {
	primitive.ObjectID | uint | ~string
}

type Entity[T EntityIDAdapter] struct {
	ID        T         `bson:"_id" json:"id" gorm:"primarykey"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func (b *Entity[T]) GetID() any {
	return b.ID
}

func (b *Entity[T]) ConvertIDToString() string {
	switch id := b.GetID().(type) {
	case primitive.ObjectID:
		return id.Hex()
	default:
		return b.GetID().(string)
	}
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
