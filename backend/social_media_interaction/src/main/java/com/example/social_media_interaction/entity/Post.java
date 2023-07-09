package com.example.social_media_interaction.entity;

import java.sql.Date;
import java.util.List;

public class Post {
    private Long Id;
    private Content PostContent;
    private Long UserId;
    private List<Long> LikedBy;
}
