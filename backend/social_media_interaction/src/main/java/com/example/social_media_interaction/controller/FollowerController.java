package com.example.social_media_interaction.controller;

import com.example.social_media_interaction.entity.Followers;
import com.example.social_media_interaction.service.FollowerService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/v1")
public class FollowerController {

    @Autowired
    private FollowerService followerService;

    @GetMapping("/follower/{id}")
    public ResponseEntity<Object> addFollower(@RequestParam Long id, @RequestBody Followers followers){
        return null;
    }

    @DeleteMapping("/follower/{id}")
    public ResponseEntity<Object> removeFollower(@RequestParam Long id, @RequestBody Followers followers){
        return null;
    }
}
