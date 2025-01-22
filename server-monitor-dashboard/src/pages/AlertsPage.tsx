import React from 'react';
import { Bell, AlertTriangle, CheckCircle } from 'lucide-react';

export const AlertsPage = () => {
  const mockAlerts = [
    {
      id: 1,
      type: 'critical',
      message: 'High CPU usage on prod-app-1',
      timestamp: '2024-03-15T10:30:00Z',
      status: 'active',
    },
    {
      id: 2,
      type: 'warning',
      message: 'Memory usage above 80% on prod-db-1',
      timestamp: '2024-03-15T09:15:00Z',
      status: 'resolved',
    },
  ];

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-900">Alerts</h1>
        <button className="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700">
          Configure Alerts
        </button>
      </div>

      <div className="bg-white rounded-lg shadow">
        <div className="p-6">
          <div className="flex items-center space-x-2 mb-4">
            <Bell className="w-5 h-5 text-indigo-600" />
            <h2 className="text-lg font-semibold">Recent Alerts</h2>
          </div>

          <div className="space-y-4">
            {mockAlerts.map((alert) => (
              <div
                key={alert.id}
                className={`border rounded-lg p-4 ${
                  alert.status === 'active'
                    ? 'border-red-200 bg-red-50'
                    : 'border-gray-200'
                }`}
              >
                <div className="flex items-start justify-between">
                  <div className="flex items-start space-x-3">
                    {alert.status === 'active' ? (
                      <AlertTriangle className="w-5 h-5 text-red-600 mt-1" />
                    ) : (
                      <CheckCircle className="w-5 h-5 text-green-600 mt-1" />
                    )}
                    <div>
                      <p className="font-medium">{alert.message}</p>
                      <p className="text-sm text-gray-500">
                        {new Date(alert.timestamp).toLocaleString()}
                      </p>
                    </div>
                  </div>
                  <span
                    className={`px-2 py-1 text-xs font-medium rounded-full ${
                      alert.status === 'active'
                        ? 'bg-red-100 text-red-800'
                        : 'bg-green-100 text-green-800'
                    }`}
                  >
                    {alert.status}
                  </span>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};