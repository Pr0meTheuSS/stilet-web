import React, { useState } from 'react';
import Split from 'react-split';
import DronesList from '../Components/DronesList';
import MapComponent from '../Components/Map';

// const drones = [
//     { id: 1, name: 'Drone 1', status: 'Active', longitude: 83.08970642089844, latitude: 54.84310531616211 },
//     { id: 2, name: 'Drone 2', status: 'Inactive', longitude: 83.0, latitude: 54.8 },
//     // Добавьте больше дронов с их координатами, если нужно
//   ];

const DroneControlPage = () => {
  const [selectedDrone, setSelectedDrone] = useState(null);

  return (
    <Split className="split" sizes={[30, 70]} minSize={200} gutterSize={10}>
      <DronesList onSelectDrone={setSelectedDrone} />
      <MapComponent selectedDrone={selectedDrone} />
    </Split>
  );
};

export default DroneControlPage;
