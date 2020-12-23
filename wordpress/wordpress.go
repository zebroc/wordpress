package wordpress

import (
	"fmt"
	gowp "github.com/sogko/go-wordpress"
	"log"
	"sort"
)

var (
	client       = gowp.NewClient(&gowp.Options{BaseAPIURL: WpApiBaseUrl})
	blogBaseUrl  = "https://www.zebroc.de"
	WpApiBaseUrl = blogBaseUrl + "/wp-json/wp/v2"
	perPage      = 100
)

func PrintLatestPost() {
	log.Println(findLastID(getAllPosts()))
}

func PrintBlogArticles() {
	allPosts := getAllPosts()

	log.Println("Fetched " + fmt.Sprint(len(allPosts)) + " posts:")
	log.Println("--------------------------------------------------------------------------------")
	for _, post := range allPosts {
		log.Println(fmt.Sprint(post.ID) +
			" - " + post.Date +
			": " + post.Title.Rendered +
			" (" + blogBaseUrl + "/" + post.Slug + ")")
	}
	log.Println("--------------------------------------------------------------------------------")
}

func getAllPosts() []gowp.Post {
	page := 1
	var allPosts []gowp.Post
	for {
		newPosts, _, _, err :=
			client.Posts().List(fmt.Sprintf("per_page=%d&page=%d", perPage, page))
		if err != nil {
			log.Println("Could not get posts for page " + fmt.Sprint(page))
		}
		allPosts = append(allPosts, newPosts...)
		page++
		if len(newPosts) < perPage {
			break
		}
	}

	return allPosts
}

func createPostMap(posts []gowp.Post) map[int]gowp.Post {
	m := make(map[int]gowp.Post)
	for _, post := range posts {
		m[post.ID] = post
	}
	return m
}

func findLastID(posts []gowp.Post) int {
	var ids []int

	for _, post := range posts {
		ids = append(ids, post.ID)
	}

	sort.Ints(ids)

	if len(ids) == 0 {
		return 0
	}

	return ids[len(ids)-1]
}
