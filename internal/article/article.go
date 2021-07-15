package article

import (
	contentpb "github.com/ryandpardey/grpc-demo/golang/content/v2"
)

// Article represents a news article.
type Article struct {
	// Required.
	// Unique identifier for an article.
	ID string `bson:"id"`

	// A list of authors for an article.
	Authors []*Author `bson:"authors"`

	// HTML content for an article.
	Content string `bson:"content"`

	// Images to embed in the article content.
	Images []*Image `bson:"images"`
}


// Author is an article author.
type Author struct {
	// Required.
	// Unique identifier for an author.
	ID string `bson:"id"`

	// Name is the author's name as displayed on the article byline.
	Name string `bson:"name"`

	// Email is an author's email address.
	Email string `bson:"email"`

	// Image is an author's profile image.
	Image *Image `bson:"image"`
}

// Image is a representation of an image.
type Image struct {
	// Required.
	// Unique identifier for an author.
	ID string `bson:"id"`

	// Required.
	// Must be one of ["image/jpeg", "image/gif", "image/png"].
	// Media type indicates the file format.
	MediaType string `bson:"mediaType"`

	// Text describing image content.
	Description string `bson:"description"`

	// Height in pixels.
	Height int `bson:"height"`

	// Width in pixels.
	Width int `bson:"width"`
}


// ArticleFromProto maps a protobuf Article document to a native Article document.
func ArticleFromProto(ab *contentpb.Article) *Article {
	if ab == nil {
		return nil
	}
	return &Article{
		ID:      ab.Id,
		Authors: AuthorsFromProto(ab.Authors),
		Content: ab.Content,
		Images:  ImagesFromProto(ab.Images),
	}
}

// Article FromProto maps a protobuf  document to a native Author document.
func (a *Article) ToProto() *contentpb.Article {
	if a == nil {
		return nil
	}
	return &contentpb.Article{
		Id: a.ID,
		Authors: AuthorsToProto(a.Authors),
		Content: a.Content,
		Images: ImagesToProto(a.Images),
	}
}

// AuthorsFromProto converts a slice of protobuf Author to a slice of native Author.
func AuthorsFromProto(a []*contentpb.Author) []*Author {
	var authors []*Author
	if len(a) == 0 {
		return authors
	}
	for _, j := range a {
	aa := &Author{}
	aa.FromProto(j)
		authors = append(authors, aa)
	}

	return authors
}

// AuthorsToProto converts a slice of native Author to a slice of protobuf Author.
func AuthorsToProto(a []*Author) []*contentpb.Author {
	var authors []*contentpb.Author
	if len(a) == 0 {
		return authors
	}
	for _, j := range a {
		authors = append(authors, j.ToProto())
	}

	return authors
}

// Author FromProto maps a protobuf Author document to a native Author document.
func (a *Author) FromProto(ab *contentpb.Author) {
	if ab == nil {
		return
	}
	a.ID = ab.Id
	a.Name = ab.Name
	a.Email = ab.Email
	a.Image	= ImageFromProto(ab.Image)
}

// Author ToProto maps a native Author document to a protobuf Author document.
func (a *Author) ToProto() *contentpb.Author {
	if a == nil {
		return nil
	}
	return &contentpb.Author{
		Id:    a.ID,
		Name: a.Name,
		Email: a.Email,
		Image: a.Image.ToProto(),
	}
}

// ImagesFromProto converts a slice of protobuf Image to a slice of native Image.
func ImagesFromProto(i []*contentpb.Image) []*Image {
	var images []*Image
	if len(i) == 0 {
		return images
	}
	for _, j := range i {
		images = append(images, ImageFromProto(j))
	}

	return images
}

// ImagesToProto converts a slice of native Image to a slice of protobuf Image.
func ImagesToProto(a []*Image) []*contentpb.Image {
	var images []*contentpb.Image
	if len(a) == 0 {
		return images
	}
	for _, j := range a {
		images = append(images, j.ToProto())
	}

	return images
}

// Image FromProto maps a protobuf Image document to a native Image document.
func ImageFromProto(ab *contentpb.Image) *Image {
	if ab == nil {
		return nil
	}
	return &Image{
		ID: ab.Id,
		Description: ab.Description,
		Height:  int(ab.Height),
		MediaType: ab.MediaType,
		Width: int(ab.Width),
	}
}

// Image ToProto maps a native Image document to a protobuf Image document.
func (i *Image) ToProto() *contentpb.Image {
	if i == nil {
		return nil
	}
	return &contentpb.Image{
		Id:    i.ID,
		Description: i.Description,
		Height: int32(i.Height),
		MediaType: i.MediaType,
		Width: int32(i.Width),
	}
}
