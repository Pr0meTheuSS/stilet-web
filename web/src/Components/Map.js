// MapComponent.js
import React, { useEffect, useRef } from 'react';
import maplibregl from 'maplibre-gl';
import 'maplibre-gl/dist/maplibre-gl.css'; // Добавляем стили карты

const MapComponent = ({ selectedDrone, drones }) => {
  const mapContainerRef = useRef(null);
  const mapRef = useRef(null);
  const markerRefs = useRef([]);

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

    if (drones) {
      drones.forEach(drone => {
        const { longitude, latitude } = drone;
        const marker = new maplibregl.Marker()
          .setLngLat([longitude, latitude])
          .addTo(mapRef.current)
          .getElement();

        marker.addEventListener('click', () => {
          alert(`Дрон: ${drone.name}\nСтатус: ${drone.status}`);
        });

        markerRefs.current.push(marker);
      });
    }

    return () => {
      markerRefs.current.forEach(marker => {
        if (marker) {
          marker.remove(); 
        }
      });
      if (mapRef.current) {
        mapRef.current.remove();
      }
    };
  }, [drones]);

  useEffect(() => {
    if (selectedDrone && mapRef.current) {
      const { longitude, latitude } = selectedDrone;
      const marker = new maplibregl.Marker()
        .setLngLat([longitude, latitude])
        .addTo(mapRef.current);
      
      markerRefs.current.push(marker);
      mapRef.current.setCenter([longitude, latitude]);

      return () => {
        marker.remove();
      };
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
