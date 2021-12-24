package goreddit

import "github.com/google/uuid"

// Thread
type Thread struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
}

// Post
type Post struct {
	ID       uuid.UUID `db:"id"`
	ThreadID uuid.UUID `db:"thread_id"`
	Title    string    `dd:"title"`
	Content  string    `db:"content"`
	Vote     int       `db:"vote"`
}

// Comment
type Comment struct {
	ID      uuid.UUID `db:"id"`
	PostID  uuid.UUID `db:"post_id"`
	Comment string    `db:"comment"`
	Post    int       `db:"post"`
}
