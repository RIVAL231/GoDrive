package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a user in the system

type User struct {
    // ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Username string             `json:"username"`
    Password string             `json:"password"` // hashed
}

type File struct{
	// ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	Size     int64  `json:"size"`
}