package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// The omitempty means that if there is no data in the particular field,
// when saved to MongoDB the field will not exist on the document rather than existing with an empty value.
type Article struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Author      string             `bson:"author,omitempty"`
	Link        string             `bson:"link,omitempty"`
	Description string             `bson:"description,omitempty"`
	Images      []string           `bson:"images,omitempty"`
}
