package com.example.social_media_interaction.controller;

import com.example.social_media_interaction.entity.Comment;
import com.example.social_media_interaction.service.CommentService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/v1")
public class CommentController {

    @Autowired
    private CommentService commentService;

    @GetMapping("/comments/{id}")
    public ResponseEntity<Object> getCommentById(@RequestParam Long Id){
        return null;
    }

    @PostMapping("/comments/")
    public  ResponseEntity<Object> postComment(@RequestBody Comment comment){
        return null;
    }

    @PutMapping("/comments/{id}")
    public ResponseEntity<Object> updateComment(@RequestParam Long id, @RequestBody Comment comment){
        return null;
    }

    @DeleteMapping("/comments/{id}")
    public ResponseEntity<Object> deleteComment(@RequestParam Long id){
        return null;
    }
}
