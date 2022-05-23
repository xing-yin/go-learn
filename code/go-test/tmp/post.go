package person

/**
 * @Author: Alan Yin
 * @Date: 2022/5/19
 * @Desc:
 */

type Post struct {
	Name    string
	Address string
}

type Service interface {
	listPosts() ([]*Post, error)
}

func listPosts(svc Service) ([]*Post, error) {
	return svc.listPosts()
}
