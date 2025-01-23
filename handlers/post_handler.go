package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/geuun/goroutine-example/models/common"
	post "github.com/geuun/goroutine-example/models/post"
	"io"
	"net/http"
)

const baseUrl = "https://jsonplaceholder.typicode.com"

func GetPost(rw http.ResponseWriter, r *http.Request, id int) {
	url := fmt.Sprintf("%s/posts/%d", baseUrl, id)
	resp, err := http.Get(url)
	if err != nil {
		common.SendResponse(rw, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		common.SendResponse(rw, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var pres post.Response
	err = json.Unmarshal(body, &pres)
	if err != nil {
		common.SendResponse(rw, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// response to entity
	postEntity := post.FromResponse(&pres)
	fmt.Println(postEntity)

	common.SendResponse(rw, http.StatusOK, "Post retrieved successfully", pres)
}

func GetPosts(rw http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(baseUrl + "/posts")
	if err != nil {
		common.SendResponse(rw, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		common.SendResponse(rw, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var pres []post.Response
	err = json.Unmarshal(body, &pres)
	if err != nil {
		common.SendResponse(rw, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// response to entity
	var posts []post.Post
	for _, pre := range pres {
		posts = append(posts, post.FromResponse(&pre))
	}
	fmt.Println(posts)

	// response to entity 2
	posts2 := post.FromResponses(pres)
	fmt.Println(posts2)

	common.SendResponse(rw, http.StatusOK, "Posts retrieved successfully", pres)
}
