package handler

import (
	"encoding/json"
	"fmt"
	"github.com/geuun/goroutine-example/models"
	"github.com/geuun/goroutine-example/models/dto/post"
	"io"
	"net/http"
)

const baseUrl = "https://jsonplaceholder.typicode.com"

func GetPosts(rw http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(baseUrl + "/posts")
	if err != nil {
		response := models.NewResponse(http.StatusInternalServerError, err.Error(), nil)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(response)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		response := models.NewResponse(http.StatusInternalServerError, err.Error(), nil)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(response)
		return
	}

	var posts []post.PostDto
	err = json.Unmarshal(body, &posts)
	if err != nil {
		response := models.NewResponse(http.StatusInternalServerError, err.Error(), nil)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(response)
		return
	}

	response := models.NewResponse(http.StatusOK, "Posts retrieved successfully", posts)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response)
}

func GetPost(rw http.ResponseWriter, r *http.Request, id int) {
	url := fmt.Sprintf("%s/posts/%d", baseUrl, id)
	resp, err := http.Get(url)
	if err != nil {
		response := models.NewResponse(http.StatusInternalServerError, err.Error(), nil)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(response)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		response := models.NewResponse(http.StatusInternalServerError, err.Error(), nil)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(response)
		return
	}

	var dto post.PostDto
	err = json.Unmarshal(body, &dto)
	if err != nil {
		response := models.NewResponse(http.StatusInternalServerError, err.Error(), nil)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(response)
		return
	}

	response := models.NewResponse(http.StatusOK, "Post retrieved successfully", dto)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response)
}
