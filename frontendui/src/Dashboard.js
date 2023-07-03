import React, { useState } from 'react';

function Dashboard() {
  const [posts, setPosts] = useState([]);
  const [newPost, setNewPost] = useState('');
  const [newComments, setNewComments] = useState({});

  // Function to handle creating a new post
  const handleCreatePost = () => {
    if (newPost.trim() === '') {
      return;
    }

    const post = {
      id: Date.now(),
      content: newPost,
      likes: 0,
      comments: [],
    };

    setPosts(prevPosts => [...prevPosts, post]);
    setNewPost('');
  };

  // Function to handle adding a comment to a post
  const handleAddComment = (postId) => {
    if (newComments[postId].trim() === '') {
      return;
    }

    const updatedPosts = posts.map(post => {
      if (post.id === postId) {
        return {
          ...post,
          comments: [...post.comments, newComments[postId]],
        };
      }
      return post;
    });

    setPosts(updatedPosts);
    setNewComments(prevComments => ({
      ...prevComments,
      [postId]: ''
    }));
  };

  // Function to handle adding a like to a post
  const handleAddLike = (postId) => {
    const updatedPosts = posts.map(post => {
      if (post.id === postId) {
        return {
          ...post,
          likes: post.likes + 1,
        };
      }
      return post;
    });

    setPosts(updatedPosts);
  };

  // Function to handle following a user
  const handleFollowUser = (postId) => {
    // Logic to follow a user
  };

  // Function to handle logging out
  const handleLogout = () => {
    // Logic to logout the user
  };

  // Function to navigate to the user's profile
  const navigateToProfile = () => {
    // Logic to navigate to the user's profile
  };

  // Function to update the new comment value for a specific post
  const handleCommentChange = (postId, value) => {
    setNewComments(prevComments => ({
      ...prevComments,
      [postId]: value
    }));
  };

  return (
    <div className="App">
      <div className="header">
        <h1>Social Media Application</h1>
        <div className="buttons-container">
          <button className="profile-button" onClick={navigateToProfile}>
            Profile
          </button>
          <button className="logout-button" onClick={handleLogout}>
            Logout
          </button>
        </div>
      </div>
      <div className="post-form">
        <textarea
          className="post-input"
          placeholder="What's on your mind?"
          value={newPost}
          onChange={e => setNewPost(e.target.value)}
        />
        <button className="post-button" onClick={handleCreatePost}>
          Create Post
        </button>
      </div>
      <div className="post-list">
        {posts.map(post => (
          <div key={post.id} className="post">
            <p>{post.content}</p>
            <div className="post-actions">
              <button className="like-button" onClick={() => handleAddLike(post.id)}>
                Like ({post.likes})
              </button>
              <button className="follow-button" onClick={() => handleFollowUser(post.id)}>
                Follow
              </button>
            </div>
            <div className="comments-section">
              {post.comments.map((comment, index) => (
                <p key={index} className="comment">{comment}</p>
              ))}
            </div>
            <div className="add-comment-section">
              <input
                type="text"
                className="comment-input"
                placeholder="Add a comment"
                value={newComments[post.id] || ''}
                onChange={e => handleCommentChange(post.id, e.target.value)}
              />
              <button
                className="comment-button"
                onClick={() => handleAddComment(post.id)}
              >
                Add Comment
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Dashboard;
