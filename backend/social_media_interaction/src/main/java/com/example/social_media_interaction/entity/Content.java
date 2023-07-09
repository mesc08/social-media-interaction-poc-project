package com.example.social_media_interaction.entity;

import jakarta.persistence.Entity;

import java.util.Date;
import java.util.List;

public class Content {
    private String Title;
    private String Content;
    private List<String> Attachment;
    private Date CreatedAt;
    private Date UpdatedAt;
}
