import React from 'react';
import { AlertTriangle, CheckCircle2 } from 'lucide-react';
import { DashboardCard } from './DashboardCard';

export function AlertsTimeline() {
  const alerts = [
    { type: 'warning', message: 'High CPU usage detected', time: '2m ago' },
    { type: 'resolved', message: 'Memory usage normalized', time: '15m ago' },
    { type: 'warning', message: 'Network latency spike', time: '1h ago' },
    { type: 'resolved', message: 'Database backup completed', time: '2h ago' },
  ];

  return (
    <DashboardCard title="Recent Alerts">
      <div className="space-y-4">
        {alerts.map((alert, index) => (
          <div key={index} className="flex items-start space-x-3">
            {alert.type === 'warning' ? (
              <AlertTriangle className="w-5 h-5 text-yellow-500 mt-0.5" />
            ) : (
              <CheckCircle2 className="w-5 h-5 text-green-500 mt-0.5" />
            )}
            <div className="flex-1">
              <p className="text-sm font-medium text-gray-900">{alert.message}</p>
              <p className="text-xs text-gray-500">{alert.time}</p>
            </div>
          </div>
        ))}
      </div>
    </DashboardCard>
  );
}