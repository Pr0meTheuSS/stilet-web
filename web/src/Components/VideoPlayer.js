import React from 'react';

const VideoPlayer = ({ src, onClose }) => {
  return (
    <div className="video-player">
      <div className="overlay" onClick={onClose}></div>
      <div className="video-container">
        <h2>Трансляция видео</h2>
        <video controls autoPlay>
          <source src={src} type="video/mp2t" /> {/* Замените на ваш формат видео */}
          Ваш браузер не поддерживает видео.
        </video>
        <button onClick={onClose}>Закрыть</button>
      </div>
    </div>
  );
};

export default VideoPlayer;
