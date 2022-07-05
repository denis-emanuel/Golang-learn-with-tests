package blogposts

import (
	"io/fs"
)

// NewPostsFromFS returns a collection of blog posts from a file system. If it does not conform to the format then it'll return an error
// fs.FS gives us some elegant ways of reading data from the file systems
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	
	//ReadDir returns a []DirEntry
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	//create a new post for each file we encounter
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}