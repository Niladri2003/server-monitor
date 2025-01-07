import React from 'react';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import { Server, Activity } from 'lucide-react';
import 'leaflet/dist/leaflet.css';
import L from 'leaflet';

// Fix for default marker icons
delete (L.Icon.Default.prototype as any)._getIconUrl;
L.Icon.Default.mergeOptions({
  iconRetinaUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-icon-2x.png',
  iconUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-icon.png',
  shadowUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-shadow.png',
});

const serverLocations = [
  { name: 'North America', position: [40.7128, -74.0060], latency: '23ms', status: 'Excellent' },
  { name: 'Europe', position: [51.5074, -0.1278], latency: '45ms', status: 'Good' },
  { name: 'Asia Pacific', position: [35.6762, 139.6503], latency: '89ms', status: 'Fair' },
  { name: 'South America', position: [-23.5505, -46.6333], latency: '120ms', status: 'Fair' },
];

export default function GlobalMonitoring() {
  return (
    <section id="global" className="py-16 md:py-20 bg-gray-900 text-white">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center mb-12 md:mb-16">
          <h2 className="text-2xl md:text-3xl font-bold">Global Port Monitoring</h2>
          <p className="mt-4 text-lg md:text-xl text-gray-400">Real-time server performance across regions</p>
        </div>
        
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          {/* Interactive Map */}
          <div className="bg-gray-800 rounded-xl p-4 h-[300px] md:h-[500px] order-2 lg:order-1">
            <MapContainer
              center={[20, 0]}
              zoom={2}
              style={{ height: '100%', width: '100%', borderRadius: '0.75rem' }}
              zoomControl={false}
            >
              <TileLayer
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
              />
              {serverLocations.map((server) => (
                <Marker key={server.name} position={server.position as [number, number]}>
                  <Popup>
                    <div className="text-gray-900">
                      <h3 className="font-semibold">{server.name}</h3>
                      <p>Latency: {server.latency}</p>
                      <p>Status: {server.status}</p>
                    </div>
                  </Popup>
                </Marker>
              ))}
            </MapContainer>
          </div>

          {/* Real-time Metrics */}
          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 order-1 lg:order-2">
            {serverLocations.map((server) => (
              <div key={server.name} className="bg-gray-800 rounded-xl p-4 md:p-6 relative overflow-hidden">
                <div className="absolute inset-0 bg-gradient-to-br from-indigo-500/10 to-purple-500/10"></div>
                <div className="relative z-10">
                  <div className="flex justify-between items-center mb-4">
                    <div className="flex items-center space-x-2">
                      <Server className="w-5 h-5 text-indigo-400" />
                      <span className="text-gray-200 text-sm md:text-base">{server.name}</span>
                    </div>
                    <Activity className="w-4 h-4 text-green-400" />
                  </div>
                  <div className="flex justify-between items-end">
                    <div>
                      <div className="text-xs md:text-sm text-gray-400 mb-1">Latency</div>
                      <div className="text-xl md:text-2xl font-bold">{server.latency}</div>
                    </div>
                    <div className="text-xs md:text-sm text-gray-400">{server.status}</div>
                  </div>
                  <div className="w-full bg-gray-700 rounded-full h-2 mt-4">
                    <div 
                      className="bg-indigo-500 rounded-full h-2 transition-all duration-500"
                      style={{ 
                        width: `${100 - parseInt(server.latency)}%`
                      }}
                    ></div>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </section>
  );
}