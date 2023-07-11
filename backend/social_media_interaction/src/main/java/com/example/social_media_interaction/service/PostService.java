package com.example.social_media_interaction.service;

import com.example.social_media_interaction.entity.Post;
import com.example.social_media_interaction.repository.PostRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class PostService {

    @Autowired
    private PostRepository postRepository;

    private void getAllPost(Long userid){

    }

    private void getPostById(Long postid){

    }

    private void addPost(Post post){

    }

    private void updatePost(Post post, Long postid){

    }

    private void deletePost(Long postid){

    }
}
