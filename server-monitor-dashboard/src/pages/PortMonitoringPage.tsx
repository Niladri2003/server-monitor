import React from 'react';
import { Activity } from 'lucide-react';
import { RequestMap } from '../components/maps/RequestMap';
import { StatsCard } from '../components/monitoring/StatsCard';
import { RequestLog } from '../components/monitoring/RequestLog';
import { RequestLocation, MonitoringStats } from '../types/monitoring';

const mockLocations: RequestLocation[] = [
  { name: "New York", coordinates: [-74.006, 40.7128], requests: 1500 },
  { name: "London", coordinates: [-0.1276, 51.5074], requests: 1200 },
  { name: "Tokyo", coordinates: [139.6917, 35.6895], requests: 2000 },
  { name: "Sydney", coordinates: [151.2093, -33.8688], requests: 800 },
  { name: "Singapore", coordinates: [103.8198, 1.3521], requests: 1700 },
  { name: "Mumbai", coordinates: [72.8777, 19.0760], requests: 900 },
  { name: "São Paulo", coordinates: [-46.6333, -23.5505], requests: 600 },
];

const mockStats: MonitoringStats = {
  totalRequests: 8700,
  activePorts: 24,
  avgResponseTime: 142,
  failedRequests: 23,
};

export const PortMonitoringPage = () => {
  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-900">Port Monitoring</h1>
        <div className="flex items-center space-x-2 text-sm text-gray-600">
          <Activity className="w-4 h-4" />
          <span>Live Updates</span>
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-4 gap-6">
        <StatsCard
          title="Total Requests"
          value={mockStats.totalRequests}
          trend={{ value: "12% from last hour", direction: "up" }}
        />
        <StatsCard
          title="Active Ports"
          value={mockStats.activePorts}
          subtitle="Across 7 locations"
        />
        <StatsCard
          title="Avg Response Time"
          value={`${mockStats.avgResponseTime}ms`}
          trend={{ value: "8% from last hour", direction: "down" }}
        />
        <StatsCard
          title="Failed Requests"
          value={mockStats.failedRequests}
          trend={{ value: "2% from last hour", direction: "up" }}
        />
      </div>

      <div className="bg-white p-6 rounded-lg shadow-md">
        <RequestMap locations={mockLocations} />
      </div>

      <RequestLog locations={mockLocations} />
    </div>
  );
};