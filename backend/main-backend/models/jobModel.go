package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Job struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Title       string             `json:"title"`
    Description string             `json:"description"`
    EmployerID  primitive.ObjectID `json:"employer_id"`
    CreatedAt   int64              `json:"created_at"`
    UpdatedAt   int64              `json:"updated_at,omitempty"`
}