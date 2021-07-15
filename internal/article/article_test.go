package article

import (
	"fmt"
	"encoding/json"
	contentpb "github.com/ryandpardey/grpc-demo/golang/content/v2"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_ArticleToProto(t *testing.T) {
	cases := []struct {
		name    string
		id 		string
		article *Article
		print   bool
	}{
		{"Success (complete): article-1", "article-1", testArticleNative(t, "article-1"), false},
		{"Success (no content): article-2", "article-2", testArticleNative(t, "article-2"), true},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			pbarticle := test.article.ToProto()
			assert.Equal(t, test.article.ID, pbarticle.Id)
			assert.Equal(t, test.article.Content, pbarticle.Content)
			if test.print {
				printStruct(pbarticle)
			}
		})
	}
}

func Test_ArticleFromProto(t *testing.T) {
	cases := []struct {
		name    string
		id 		string
		article *contentpb.Article
		print   bool
	}{
		{"Success (complete): article-1", "article-1", testArticleProto(t, "article-1"), false},
		{"Success (no content): article-2", "article-2", testArticleProto(t, "article-2"), false},
		{"Success (no authors): article-3", "article-3", testArticleProto(t, "article-3"), false},
		{"Success (no images): article-4", "article-4", testArticleProto(t, "article-4"), false},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			article := ArticleFromProto(test.article)
			assert.Equal(t, test.article.Id, article.ID)
			assert.Equal(t, test.article.Content, article.Content)
			if strings.Contains(test.name, "no authors") {
				assert.Nil(t, article.Authors)
			}
			if strings.Contains(test.name, "no images") {
				assert.Nil(t, article.Images)
			}
			if test.print {
				printStruct(article)
			}
		})
	}
}

func testArticleNative(t *testing.T, a string) *Article {
	t.Helper()
	var authors []*Author
	var images []*Image
	switch a {
	case "article-1":
		authors = append(authors, testAuthorNative(t, "john-smith"), testAuthorNative(t, "jane-doe"))
		images = append(images, testImageNative(t, "image-1"), testImageNative(t, "image-2"))
		return &Article{
			ID:      a,
			Authors: authors,
			Content: "this is some article content",
			Images:  images,
		}
	case "article-2":
		authors = append(authors, testAuthorNative(t, "john-smith"), testAuthorNative(t, "jane-doe"))
		images = append(images, testImageNative(t, "image-1"), testImageNative(t, "image-2"))
		return &Article{
			ID:      a,
			Authors: authors,
			Images:  images,
		}
	}

	return nil
}

func testArticleProto(t *testing.T, a string) *contentpb.Article {
	t.Helper()
	var authors []*contentpb.Author
	var images []*contentpb.Image
	switch a {
	case "article-1":
		authors = append(authors, testAuthorProto(t, "john-smith"), testAuthorProto(t, "jane-doe"))
		images = append(images, testImageProto(t, "image-1"), testImageProto(t, "image-2"))
		return &contentpb.Article{
			Id:      a,
			Authors: authors,
			Content: "this is some article content",
			Images:  images,
		}
	case "article-2":
		authors = append(authors, testAuthorProto(t, "john-smith"), testAuthorProto(t, "jane-doe"))
		images = append(images, testImageProto(t, "image-1"), testImageProto(t, "image-2"))
		return &contentpb.Article{
			Id:      a,
			Authors: authors,
			Images:  images,
		}
	case "article-3":
		images = append(images, testImageProto(t, "image-1"), testImageProto(t, "image-2"))
		return &contentpb.Article{
			Id:      a,
			Content: "this is some article content",
			Images:  images,
		}
	case "article-4":
		authors = append(authors, testAuthorProto(t, "john-smith"), testAuthorProto(t, "jane-doe"))
		return &contentpb.Article{
			Id:      a,
			Authors: authors,
			Content: "this is some article content",
		}
	}

	return nil
}

func testAuthorNative(t *testing.T, a string) *Author {
	t.Helper()
	switch a {
	case "john-smith":
		return &Author{
			ID:    a,
			Name:  "John Smith",
			Email: "johnsmith@test.com",
			Image: testImageNative(t, a),
		}
	case "jane-doe":
		return &Author{
			ID:    a,
			Name:  "John Smith",
			Email: "johnsmith@test.com",
			Image: testImageNative(t, a),
		}
	}

	return nil
}

func testAuthorProto(t *testing.T, a string) *contentpb.Author {
	t.Helper()
	switch a {
	case "john-smith":
		return &contentpb.Author{
			Id:    a,
			Name:  "John Smith",
			Email: "johnsmith@test.com",
			Image: testImageProto(t, a),
		}
	case "jane-doe":
		return &contentpb.Author{
			Id:    a,
			Name:  "John Smith",
			Email: "johnsmith@test.com",
			Image: testImageProto(t, a),
		}
	}

	return nil
}


func testImageNative(t *testing.T, i string) *Image {
	t.Helper()
	switch i{
	case "john-smith":
		return &Image{
			ID: "john-smith-image",
			Description: "this is a test image for john smith",
			Height: 400,
			MediaType: "image/png",
			Width: 800,
		}
	case "jane-doe":
		return &Image{
			ID: "jane-doe-image",
			Description: "this is a test image for jane doe",
			Height: 500,
			MediaType: "image/jpeg",
			Width: 1080,
		}
	case "image-1":
		return &Image{
			ID: "st-image",
			Description: "this is an article test image",
			Height: 1200,
			MediaType: "image/gif",
			Width: 4200,
		}
	case "image-2":
		return &Image{
			ID: "test-image",
			Description: "this is another article test image",
			Height: 720,
			MediaType: "image/png",
			Width: 1350,
		}
	}

	return nil
}

func testImageProto(t *testing.T, i string) *contentpb.Image {
	t.Helper()
	switch i{
	case "john-smith":
		return &contentpb.Image{
			Id: "john-smith-image",
			Description: "this is a test image for john smith",
			Height: 400,
			MediaType: "image/png",
			Width: 800,
		}
	case "jane-doe":
		return &contentpb.Image{
			Id: "jane-doe-image",
			Description: "this is a test image for jane doe",
			Height: 500,
			MediaType: "image/jpeg",
			Width: 1080,
		}
	case "image-1":
		return &contentpb.Image{
			Id: "st-image",
			Description: "this is an article test image",
			Height: 1200,
			MediaType: "image/gif",
			Width: 4200,
		}
	case "image-2":
		return &contentpb.Image{
			Id: "test-image",
			Description: "this is another article test image",
			Height: 720,
			MediaType: "image/png",
			Width: 1350,
		}
	}

	return nil
}

func printStruct(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}
