package posts

type Getter interface {
	GetAll() []Post
}

type Adder interface {
	Add(post Post)
}

type Post struct {
	Title string `json:"title"`
	Text  string `json:"post"`
}

type Repository struct {
	Posts []Post
}

func New() *Repository {
	return &Repository{
		Posts: []Post{},
	}
}

func (r *Repository) GetAll() []Post {
	return r.Posts
}

func (r *Repository) Add(post Post) {
	r.Posts = append(r.Posts, post)
}
