import React from 'react';
import './Profile.css';

function Profile() {
  // Sample user data
  const user = {
    name: 'John Doe',
    username: 'johndoe',
    followers: 500,
    following: 250,
    photos: [
      'photo1.jpg',
      'photo2.jpg',
      'photo3.jpg',
      'photo4.jpg',
    ],
  };

  // Function to handle clicking the back button
  const handleBack = () => {
    // Logic to perform any necessary actions before redirecting
    window.location.href = '/dashboard';
  };

  return (
    <div className="profile-container">
      <h2 className="profile-heading">Profile</h2>
      {/* User details */}
      <div className="user-details">
        <h3 className="section-heading">User Details</h3>
        <p>
          <strong>Name:</strong> {user.name}
        </p>
        <p>
          <strong>Username:</strong> {user.username}
        </p>
      </div>
      {/* Photos */}
      <div className="photos">
        <h3 className="section-heading">Photos</h3>
        <div className="photo-grid">
          {user.photos.map((photo, index) => (
            <img
              key={index}
              className="photo"
              src={photo}
              alt={`Photo ${index + 1}`}
            />
          ))}
        </div>
      </div>
      {/* Followers */}
      <div className="followers">
        <h3 className="section-heading">Followers</h3>
        <p>{user.followers}</p>
      </div>
      {/* Following */}
      <div className="following">
        <h3 className="section-heading">Following</h3>
        <p>{user.following}</p>
      </div>
      {/* Back button */}
      <button className="back-button" onClick={handleBack}>
        Back
      </button>
    </div>
  );
}

export default Profile;
