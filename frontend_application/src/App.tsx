import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css'; // Import CSS

const App: React.FC = () => {
  const [content, setContent] = useState<string>(''); // Notification input state
  const [notification, setNotification] = useState<string | null>(null); // Real-time notification state

  // Handle manual notification submission
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const result = await axios.post(
        'http://localhost:8080/submit',
        new URLSearchParams({ content }),
        { headers: { 'Content-Type': 'application/x-www-form-urlencoded' } }
      );
      setNotification(result.data); // Show notification as popup
      setContent(''); // Clear input field
    } catch (error) {
      console.error(error);
      setNotification('Error submitting notification');
    }
  };

  // WebSocket connection for real-time notifications
  useEffect(() => {
    const socket = new WebSocket('ws://localhost:8080/ws'); // Connect to WebSocket

    socket.onmessage = (event) => {
      setNotification(event.data); // Display real-time notification popup
    };

    return () => socket.close(); // Cleanup WebSocket connection on unmount
  }, []);

  // Automatically clear popup after 5 seconds
  useEffect(() => {
    const timer = setTimeout(() => setNotification(null), 5000);
    return () => clearTimeout(timer);
  }, [notification]);

  return (
    <div className="App">
      <div className="content-section">
        <h2>Send a Custom Notification</h2>
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            value={content}
            onChange={(e) => setContent(e.target.value)}
            placeholder="Enter your notification"
            className="input-field"
          />
          <button type="submit">Send</button>
        </form>
      </div>

      <div className="popup-section">
        <div className="banner">Real-Time Notification App</div>

        {/* Device screen mockup for notification display */}
        <div className="device-screen">
          {notification && (
            <div className="notification-popup">
              <p>{notification}</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default App;
