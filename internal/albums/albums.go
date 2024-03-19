package albums

type Album struct {
	ID      string
	Users   []string
	Name    string
	Content []string
	Tags    []string
}

func NewAlbum() *Album {
	return new(Album)
}

func SaveAlbum() {}

func DeleteAlbum() {}
