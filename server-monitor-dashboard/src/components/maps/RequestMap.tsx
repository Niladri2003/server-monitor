import React from 'react';
import { MapContainer, TileLayer, CircleMarker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import { RequestLocation } from '../../types/monitoring';

interface RequestMapProps {
  locations: RequestLocation[];
}

export const RequestMap: React.FC<RequestMapProps> = ({ locations }) => {
  return (
    <MapContainer
      center={[20, 0]}
      zoom={2}
      style={{ height: '600px', width: '100%', borderRadius: '0.5rem' }}
      className="z-0"
    >
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      />
      {locations.map((location) => (
        <CircleMarker
          key={location.name}
          center={[location.coordinates[1], location.coordinates[0]]}
          radius={Math.sqrt(location.requests) / 4}
          fillColor="#4f46e5"
          color="#4338ca"
          weight={2}
          opacity={0.8}
          fillOpacity={0.4}
        >
          <Popup>
            <div className="text-sm">
              <p className="font-semibold">{location.name}</p>
              <p>{location.requests} requests</p>
            </div>
          </Popup>
        </CircleMarker>
      ))}
    </MapContainer>
  );
};