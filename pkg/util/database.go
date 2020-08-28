package util

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"net/url"
)

const linkCollection = "links"

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
		log.Fatal(err)
	}
	return Firestore{client, ctx}
}

// Link represents the link
type Link struct {
	URL          string
	ID           string
	GeneratedURL string
}

// SaveLink saves link information to Firestore
func (fs *Firestore) SaveLink(link Link) {
	log.Printf("Saving link: %s", link)

	_, err := fs.DB.Collection(linkCollection).Doc(link.ID).Set(fs.Ctx, map[string]interface{}{
		"URL":          url.QueryEscape(link.URL),
		"ID":           link.ID,
		"GeneratedURL": url.QueryEscape(link.GeneratedURL),
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
}

// ReadLink retrieves link information from Firestore
func (fs *Firestore) ReadLink(id string, link *Link) error {
	dsnap, err := fs.DB.Collection(linkCollection).Doc(id).Get(fs.Ctx)
	if err != nil {
		return err
	}

	if err = dsnap.DataTo(&link); err != nil {
		return err
	}
	return nil
}
