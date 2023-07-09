package com.example.social_media_interaction.entity;

import jakarta.persistence.Entity;

@Entity
public class Comment {
    private Long Id;
    private Long Postid;
    private Long Userid;
    private Content CommentContent;
}
