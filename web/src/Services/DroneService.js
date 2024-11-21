import React, { useEffect, useState, useContext, createContext } from 'react';

// Создание контекста для предоставления данных о дронах другим компонентам
const DroneContext = createContext();

// Хук для использования данных о дронах в других компонентах
export const useDrones = () => {
  return useContext(DroneContext);
};

// Компонент-сервис для управления WebSocket соединением и получения данных о дронах
export const DroneServiceProvider = ({ children }) => {
  const [drones, setDrones] = useState([]);
  const [error, setError] = useState(null);
  const [socket, setSocket] = useState(null);
  const [isConnected, setIsConnected] = useState(false);

  useEffect(() => {
    // Устанавливаем WebSocket соединение
    const ws = new WebSocket('ws://localhost:8080/ws');

    const handleOpen = () => {
      console.log('WebSocket connection established');
      setIsConnected(true);
      ws.send(JSON.stringify({ action: 'get_all_drones' }));
    };

    const handleMessage = (event) => {
      try {
        const data = JSON.parse(event.data);

        if (data.action === 'get_all_drones') {
          setDrones(data.data);
        } else if (data.action === 'drone_updated') {
          setDrones((prevDrones) =>
            prevDrones.map((drone) =>
              drone.id === data.data.id ? { ...drone, ...data.data } : drone
            )
          );
        }
      } catch (err) {
        setError('Failed to parse incoming message');
        console.error(err);
      }
    };

    const handleError = (err) => {
      setError('WebSocket error');
      console.error(err);
    };

    const handleClose = () => {
      setIsConnected(false);
      setError('WebSocket connection closed');
      console.log('WebSocket connection closed');
    };

    ws.addEventListener('open', handleOpen);
    ws.addEventListener('message', handleMessage);
    ws.addEventListener('error', handleError);
    ws.addEventListener('close', handleClose);

    setSocket(ws);

    return () => {
      ws.removeEventListener('open', handleOpen);
      ws.removeEventListener('message', handleMessage);
      ws.removeEventListener('error', handleError);
      ws.removeEventListener('close', handleClose);
      ws.close();
    };
  }, []);

  const sendMessage = (message) => {
    if (socket && isConnected) {
      socket.send(JSON.stringify(message));
    } else {
      console.log('WebSocket is not open. Message not sent:', message);
    }
  };

  return (
    <DroneContext.Provider value={{ drones, error, sendMessage, isConnected }}>
      {children}
    </DroneContext.Provider>
  );
};
