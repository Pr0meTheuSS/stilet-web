import React, { useEffect, useRef } from 'react';
import maplibregl from 'maplibre-gl';
import 'maplibre-gl/dist/maplibre-gl.css'; // Добавляем стили карты
import { useDrones } from '../Services/DroneService';

const MapComponent = ({ selectedDrone }) => {
  const { drones } = useDrones(); // Получаем данные о дронах
  const mapContainerRef = useRef(null);
  const mapRef = useRef(null);
  const markerMap = useRef(new Map()); // Map для управления маркерами

  // Инициализация карты
  useEffect(() => {
    mapRef.current = new maplibregl.Map({
      container: mapContainerRef.current,
      style: 'https://basemaps.cartocdn.com/gl/voyager-gl-style/style.json',
      center: [0, 0],
      zoom: 2,
    });

    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          const { longitude, latitude } = position.coords;
          mapRef.current.setCenter([longitude, latitude]);
          mapRef.current.setZoom(13);
        },
        (error) => {
          console.error("Ошибка получения местоположения: ", error);
        }
      );
    } else {
      console.error("Геолокация не поддерживается вашим браузером.");
    }

    return () => {
      if (mapRef.current) {
        mapRef.current.remove();
      }
    };
  }, []);

  useEffect(() => {
    if (!mapRef.current || !drones) return;
  
    const currentMarkers = new Map();
  
    drones.forEach((drone) => {
      const { id, longitude, latitude, name, status } = drone;
  
      if (!markerMap.current.has(id)) {
        // Создаем новый маркер
        const marker = new maplibregl.Marker()
          .setLngLat([longitude, latitude])
          .addTo(mapRef.current);
  
        marker.getElement().addEventListener('click', () => {
          alert(`Дрон: ${name}\nСтатус: ${status}`);
        });
  
        markerMap.current.set(id, { marker, drone });
      } else {
        // Обновляем существующий маркер
        const markerData = markerMap.current.get(id);
        if (
          markerData.drone.longitude !== longitude ||
          markerData.drone.latitude !== latitude
        ) {
          markerData.marker.setLngLat([longitude, latitude]); // Используем сам объект маркера
        }
        markerData.drone = drone; // Обновляем данные дрона
      }
      currentMarkers.set(id, true); // Отмечаем маркер как актуальный
    });
  
    // Удаляем маркеры, которых больше нет
    markerMap.current.forEach((_, id) => {
      if (!currentMarkers.has(id)) {
        const { marker } = markerMap.current.get(id);
        marker.remove();
        markerMap.current.delete(id);
      }
    });
  }, [drones]);
  

  // Центрирование карты на выбранном дроне
  useEffect(() => {
    if (selectedDrone && mapRef.current) {
      const { longitude, latitude } = selectedDrone;
      mapRef.current.setCenter([longitude, latitude]);
    }
  }, [selectedDrone]);

  return (
    <div style={{ position: 'relative', width: '100%', height: '100%' }}>
      <div
        ref={mapContainerRef}
        style={{ width: '100%', height: '100%', border: '1px solid #ccc' }}
      />
    </div>
  );
};

export default MapComponent;
