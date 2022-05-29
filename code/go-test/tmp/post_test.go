package person

import (
	"testing"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/19
 * @Desc:
 */

type fakeService struct {
}

func NewFakeService() Service {
	return &fakeService{}
}

func (s *fakeService) listPosts() ([]*Post, error) {
	posts := make([]*Post, 0)
	posts = append(posts, &Post{
		Name:    "Alan",
		Address: "hz",
	})
	posts = append(posts, &Post{
		Name:    "Bob",
		Address: "bj",
	})
	return posts, nil
}

func Test_listPosts(t *testing.T) {
	fake := NewFakeService()
	if _, err := listPosts(fake); err != nil {
		t.Fatal("list posts failed")
	}
}
