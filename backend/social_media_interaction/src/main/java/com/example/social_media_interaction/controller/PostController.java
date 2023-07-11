package com.example.social_media_interaction.controller;

import com.example.social_media_interaction.entity.Post;
import com.example.social_media_interaction.service.PostService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/v1")
public class PostController {

    @Autowired
    private PostService postService;

    @GetMapping("/posts")
    public ResponseEntity<Object> getAllPostByFollower(){
        return null;
    }

    @GetMapping("/posts/{id}")
    public ResponseEntity<Object> getPostByIdOfFollower(@RequestParam Long id){
        return null;
    }

    @PostMapping("/posts")
    public ResponseEntity<Object> addPost(@RequestBody Post postData){
        return null;
    }

    @PutMapping("/posts/{id}")
    public ResponseEntity<Object> updatePost(@RequestBody Post postData, @RequestParam Long id){
        return null;
    }

    @DeleteMapping("/posts/{id}")
    public ResponseEntity<Object> deletePost(@RequestParam Long id){
        return null;
    }

}
