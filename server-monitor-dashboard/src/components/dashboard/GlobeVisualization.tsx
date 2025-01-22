import React, { useEffect, useRef } from 'react';
import Globe from 'react-globe.gl';

interface PortConnection {
  startLat: number;
  startLng: number;
  endLat: number;
  endLng: number;
  color: string;
}

const generateRandomConnections = (count: number): PortConnection[] => {
  return Array.from({ length: count }, () => ({
    startLat: (Math.random() - 0.5) * 180,
    startLng: (Math.random() - 0.5) * 360,
    endLat: (Math.random() - 0.5) * 180,
    endLng: (Math.random() - 0.5) * 360,
    color: ['#ff5733', '#33ff57', '#3357ff'][Math.floor(Math.random() * 3)],
  }));
};

export const GlobeVisualization = () => {
  const globeEl = useRef<any>();
  const connections = generateRandomConnections(20);

  useEffect(() => {
    if (globeEl.current) {
      globeEl.current.controls().autoRotate = true;
      globeEl.current.controls().autoRotateSpeed = 0.5;
    }
  }, []);

  const arcData = connections.map(conn => ({
    startLat: conn.startLat,
    startLng: conn.startLng,
    endLat: conn.endLat,
    endLng: conn.endLng,
    color: conn.color,
  }));

  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <h3 className="text-lg font-semibold mb-4">Port Traffic Visualization</h3>
      <div className="h-[500px] w-full">
        <Globe
          ref={globeEl}
          globeImageUrl="//unpkg.com/three-globe/example/img/earth-dark.jpg"
          arcsData={arcData}
          arcColor="color"
          arcDashLength={() => Math.random()}
          arcDashGap={() => Math.random()}
          arcDashAnimateTime={() => Math.random() * 4000 + 500}
          backgroundColor="rgba(0,0,0,0)"
          width={800}
          height={500}
        />
      </div>
    </div>
  );
};