package models

type Hair struct {
	Id           string   `firestore:"id"`
	Name         string   `firestore:"name"`
	NameQuery    []string `firestore:"name_query"`
	Description  string   `firestore:"description"`
	Images       string   `firestore:"images"`
	CommentCount int      `firestore:"comment_count"`
	Rating       float32  `firestore:"rating"`
}
