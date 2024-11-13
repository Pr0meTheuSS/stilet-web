import React from 'react';

const DroneInfo = ({ drone }) => {
  if (!drone) return null;

  return (
    <div className="drone-info">
      <h2>{drone.name}</h2>
      <p>Status: {drone.status}</p>
      <p>Longitude: {drone.longitude}</p>
      <p>Latitude: {drone.latitude}</p>
    </div>
  );
};

export default DroneInfo;
