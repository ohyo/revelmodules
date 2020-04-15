package blog

import (
	"github.com/ohyo/revelmodules/blog/app/models"
	"github.com/revel/revel"
)

// StorageInterface is
// use makeSlug for slugging, i.e. scale-golang-applications
type StorageInterface interface {
	// GetUser is return User based on request
	// posssible to use cache, session, database or something else
	GetUser(ctrl revel.Controller)
	// ListPosts get the list of posts for the user by slug and page
	// if no user slug so it return post ordered by create date
	ListPosts(userSlug string, page int) []*models.Post
	GetPost(userSlug string, postSlug string) *models.Post
	// SavePost is save post. If there no id so it create new post, otherwise it update it
	SavePost(post models.Post)
	// RemovePost is remove the post using userSlug and postSlug. Both params required
	RemovePost(userSlug string, postSlug string)
	// SaveComment is save comment. If there no id so it create new comment, otherwise it update it
	SaveComment(postSlug string, post models.Comment)
	// RemoveComment is remove the comment using postSlug and commentID. Both params required
	RemoveComment(postSlug string, commentID string)
}

// Storage interface
var Storage *StorageInterface
