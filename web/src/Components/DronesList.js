import React, { useState } from 'react';
import { useDrones } from '../Services/DroneService'; // Хук для получения данных о дронах
import VideoPlayer from './VideoPlayer';

const DroneList = ({ onSelectDrone }) => {
  const { drones, error } = useDrones(); // Получаем данные о дронах
  const [videoSrc, setVideoSrc] = useState(null);
  const [isVideoOpen, setIsVideoOpen] = useState(false);

  const handleVideoOpen = (drone) => {
    setVideoSrc(`rtp://example.com/${drone.id}`);
    setIsVideoOpen(true);
  };

  const handleVideoClose = () => {
    setIsVideoOpen(false);
    setVideoSrc(null);
  };

  // Обработка ошибки, если WebSocket не может получить данные
  if (error) {
    return <div>Error: {error}</div>;
  }

  // Если дронов еще нет (пока данные не получены), показываем сообщение
  if (drones.length === 0) {
    return <div>Loading drones...</div>;
  }

  return (
    <div className="drone-list">
      {drones.map((drone) => (
        <div key={drone.id} className="drone-item">
          <h3>{drone.name}</h3>
          <p>Status: {drone.status}</p>
          <p>
            Location: ({drone.longitude.toFixed(6)}, {drone.latitude.toFixed(6)})
          </p>
          <button onClick={() => handleVideoOpen(drone)}>Видео</button>
          <button onClick={() => onSelectDrone(drone)}>Отслеживать</button>
        </div>
      ))}
      {isVideoOpen && <VideoPlayer src={videoSrc} onClose={handleVideoClose} />}
    </div>
  );
};

export default DroneList;
