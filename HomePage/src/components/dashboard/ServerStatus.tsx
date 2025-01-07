import React from 'react';
import { CheckCircle, AlertCircle, Clock } from 'lucide-react';
import { DashboardCard } from './DashboardCard';

export function ServerStatus() {
  const servers = [
    { name: 'Production API', status: 'healthy', uptime: '99.99%' },
    { name: 'Database Cluster', status: 'warning', uptime: '99.95%' },
    { name: 'Cache Server', status: 'healthy', uptime: '99.99%' },
    { name: 'Load Balancer', status: 'healthy', uptime: '100%' },
  ];

  const getStatusIcon = (status: string) => {
    switch (status) {
      case 'healthy':
        return <CheckCircle className="w-5 h-5 text-green-500" />;
      case 'warning':
        return <AlertCircle className="w-5 h-5 text-yellow-500" />;
      default:
        return <Clock className="w-5 h-5 text-gray-500" />;
    }
  };

  return (
    <DashboardCard title="Server Status">
      <div className="space-y-4">
        {servers.map((server) => (
          <div key={server.name} className="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div className="flex items-center space-x-3">
              {getStatusIcon(server.status)}
              <span className="font-medium text-gray-700">{server.name}</span>
            </div>
            <span className="text-sm text-gray-500">{server.uptime}</span>
          </div>
        ))}
      </div>
    </DashboardCard>
  );
}