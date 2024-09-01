package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Application struct {
    ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    JobID      primitive.ObjectID `json:"job_id"`
    EmployeeID primitive.ObjectID `json:"employee_id"`
    AppliedAt  int64              `json:"applied_at"`
}