package utils

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

// Firestore instance
type Firestore struct {
	DB  *firestore.Client
	Ctx context.Context
}

// NewFirestore creates new instance of Firestore
func NewFirestore(projectID string) Firestore {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		// TODO: Handle error.
	}
	return Firestore{client, ctx}
}

// SaveLink saves link information to Firestore
func (fs *Firestore) SaveLink(link string, name string, shortLink string) {
	_, err := fs.DB.Collection("links").Doc("key").Set(fs.Ctx, map[string]interface{}{
		"name":      name,
		"link":      link,
		"shortLink": shortLink,
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
}
