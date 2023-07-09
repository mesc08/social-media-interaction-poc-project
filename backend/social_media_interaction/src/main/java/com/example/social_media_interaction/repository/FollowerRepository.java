package com.example.social_media_interaction.repository;

import com.example.social_media_interaction.entity.Followers;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface FollowerRepository extends JpaRepository<Followers, Long> {
}
