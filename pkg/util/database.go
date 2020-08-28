package util

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/thanhpk/randstr"
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

// Link represents link redirect information
type Link struct {
	// Original url provided by users
	URL string

	// Custom value provided as alternative name for the link redirect
	// For example https://tnij.to/{Value}
	Value string

	// Link generated using app configuration and Value
	// For example https://tnij.to/{Value}
	GeneratedURL string

	// Redirect counter
	Views int
}

// NewLink creates new Link with 0 views
func NewLink(originalURL string, id string) Link {
	if id == "" {
		id = randstr.String(11)
	}
	generatedURL := fmt.Sprintf("https://%s/%s", Config.Hostname, id)
	return Link{originalURL, id, generatedURL, 0}
}

func (l Link) escaped() Link {
	return Link{
		url.QueryEscape(l.URL),
		l.Value,
		url.QueryEscape(l.GeneratedURL),
		l.Views,
	}
}

// SaveLink saves link information to Firestore
func (fs *Firestore) SaveLink(link Link) error {
	log.Printf("Saving link: %s", link)
	_, err := fs.DB.Collection(linkCollection).Doc(link.Value).Set(fs.Ctx, link.escaped())
	return err
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

// UpdateViewsCount updates link views count
func (fs *Firestore) UpdateViewsCount(link Link) {
	var l Link
	if err := fs.ReadLink(link.Value, &l); err != nil {
		log.Printf("Updating views count failed %s", err)
	}
	_, err := fs.DB.Collection(linkCollection).Doc(link.Value).Set(fs.Ctx, map[string]interface{}{
		"Views": l.Views + 1}, firestore.MergeAll)
	if err != nil {
		log.Printf("Updating views count failed %s", err)
	}
}
