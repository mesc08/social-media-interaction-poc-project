package com.example.social_media_interaction.controller;


import com.example.social_media_interaction.entity.Followers;
import com.example.social_media_interaction.service.FollowingService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/v1")
public class FollowingController {

    @Autowired
    private FollowingService followingService;


    @GetMapping("/following/{id}")
    public ResponseEntity<Object> addFollowing(@RequestParam Long id, @RequestBody Followers followers){
        return null;
    }

    @DeleteMapping("/following/{id}")
    public ResponseEntity<Object> removeFollowing(@RequestParam Long id, @RequestBody Followers followers){
        return null;
    }
}
