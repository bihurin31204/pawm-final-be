// models/user.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct represents the data structure for MongoDB
type User struct {
    ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Username             string             `bson:"username" json:"username"`
    Password             string             `bson:"password,omitempty" json:"-"`
    LastSimulation       string             `bson:"lastSimulation,omitempty" json:"lastSimulation,omitempty"`
    CompletedSimulations []string           `bson:"completedSimulations,omitempty" json:"completedSimulations,omitempty"`
    State                map[string]interface{} `bson:"state,omitempty" json:"state,omitempty"`
}
