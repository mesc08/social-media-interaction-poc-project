package com.example.social_media_interaction.service;


import com.example.social_media_interaction.entity.Comment;
import com.example.social_media_interaction.repository.CommentRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class CommentService {

    @Autowired
    private CommentRepository commentRepository;

    private void addComment(Comment comment){}

    private Comment getComment(Long postId){
        return null;
    }

    private Comment getCommentById(Long commentId){
        return null;
    }

    private Comment updateComment(Long commentId){
        return null;
    }

    private void deleteComment(Long commentId){
    }
}
