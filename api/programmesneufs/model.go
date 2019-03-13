package programmesneufs

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

/**
how to convert ObjectID:
- declare type = primitive.ObjectID from bson/primitive
- use tag declaration bson:"_id"
- json tags only for mapping presentation
- if null field: put *string
*/
type Programmesneufs struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name                *string            `bson:"name,omitempty" json:"name,omitempty"`
	Description         *string            `json:"description,omitempty"`
	Price               interface{}        `bson:"price,omitempty" json:"price,omitempty"`
	CreationDate        string             `bson:"creationDate,omitempty" json:"createdAt,omitempty"`
	Coordinates         map[string]interface{}        `bson:"coordinates,omitempty" json:"coordinates,omitempty"`
	ProfessionalLogoUrl map[string]interface{}        `bson:"professionalLogoUrl,omitempty" json:"professionalLogoUrl,omitempty"`
	ProfessionalName 		interface{}        `bson:"professionalName,omitempty" json:"professionalName,omitempty"`
	Thumbnail           interface{}        `bson:"thumbnail,omitempty" json:"thumbnail,omitempty"`
}

/**
  id: string;
  description?: string;
  logo?: string;
  name?: string;
  onClick?: (event: MouseEvent) => void;
  title?: string;
  price?: number;
  // eslint-disable-next-line
  thumbnails?: any;
*/
