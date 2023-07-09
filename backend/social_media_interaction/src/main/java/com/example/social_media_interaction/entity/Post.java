package com.example.social_media_interaction.entity;

import jakarta.persistence.Entity;

import java.sql.Date;
import java.util.List;

@Entity
public class Post {
    private Long Id;
    private Content PostContent;
    private Long UserId;
    private List<Long> LikedBy;
    private List<Comment> Comments;
}
