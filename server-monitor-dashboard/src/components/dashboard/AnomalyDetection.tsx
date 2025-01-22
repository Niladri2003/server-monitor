import React from 'react';
import { AlertTriangle } from 'lucide-react';

interface Anomaly {
  id: string;
  type: 'warning' | 'critical';
  message: string;
  timestamp: string;
  metric: string;
}

interface AnomalyDetectionProps {
  anomalies: Anomaly[];
}

export const AnomalyDetection: React.FC<AnomalyDetectionProps> = ({ anomalies }) => {
  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <div className="flex items-center justify-between mb-4">
        <h3 className="text-lg font-semibold">Anomaly Detection</h3>
        <AlertTriangle className="w-5 h-5 text-indigo-600" />
      </div>
      <div className="space-y-3">
        {anomalies.map((anomaly) => (
          <div
            key={anomaly.id}
            className={`p-3 rounded-lg ${
              anomaly.type === 'critical' ? 'bg-red-50' : 'bg-yellow-50'
            }`}
          >
            <div className="flex items-start">
              <AlertTriangle
                className={`w-5 h-5 mt-0.5 ${
                  anomaly.type === 'critical' ? 'text-red-600' : 'text-yellow-600'
                }`}
              />
              <div className="ml-3">
                <p className="text-sm font-medium">{anomaly.message}</p>
                <div className="flex items-center mt-1 text-xs text-gray-500">
                  <span>{anomaly.metric}</span>
                  <span className="mx-2">•</span>
                  <span>{anomaly.timestamp}</span>
                </div>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};