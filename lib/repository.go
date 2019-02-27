package lib

import "fmt"

type Repository struct {
	Owner    string
	Name     string
	CloneURL string
}

func NewRepository(ownder string, name string, cloneUrl string) *Repository {
	return &Repository{
		Owner:    ownder,
		Name:     name,
		CloneURL: cloneUrl,
	}
}

func (r *Repository) FullName() string {
	return fmt.Sprintf("%s/%s", r.Owner, r.Name)
}
