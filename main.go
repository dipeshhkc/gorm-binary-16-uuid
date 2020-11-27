package main

import (
	"fmt"
	"log"
)

func main() {

	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Post{}, &Comment{})
	post := &Post{Name: "Post2"}
	if db.Create(&post).Error != nil {
		log.Panic("Unable to create post")
	}

	log.Print("post ko id", post.ID)

	db.Where("id=?", post.ID).First(&post)
	log.Print("where", post)

	comment := &Comment{Name: "Nice Post", PostID: post.ID}
	if db.Create(&comment).Error != nil {
		log.Panic("Unable to create comment")
	}
	fetchedPost := &Post{}
	if err := db.Where("id = ?", comment.PostID).Preload("Comment").First(&fetchedPost).Error; err != nil {
		log.Panic("Unable to find created Post.")
	}
	fmt.Printf("Post: %+v\n", fetchedPost)
}
