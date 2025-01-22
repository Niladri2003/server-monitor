import React from 'react';
import { Shield } from 'lucide-react';

interface SystemHealthScoreProps {
  score: number;
  metrics: {
    cpu: number;
    memory: number;
    disk: number;
    network: number;
  };
}

export const SystemHealthScore: React.FC<SystemHealthScoreProps> = ({ score, metrics }) => {
  const getHealthColor = (score: number) => {
    if (score >= 90) return 'text-green-500';
    if (score >= 70) return 'text-yellow-500';
    return 'text-red-500';
  };

  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <div className="flex items-center justify-between mb-4">
        <h3 className="text-lg font-semibold">System Health Score</h3>
        <Shield className="w-5 h-5 text-indigo-600" />
      </div>
      <div className="flex items-center justify-center mb-6">
        <div className={`text-4xl font-bold ${getHealthColor(score)}`}>
          {score}%
        </div>
      </div>
      <div className="space-y-3">
        {Object.entries(metrics).map(([key, value]) => (
          <div key={key} className="flex items-center justify-between">
            <span className="text-sm text-gray-600 capitalize">{key}</span>
            <div className="flex items-center">
              <div className="w-32 h-2 bg-gray-200 rounded-full mr-2">
                <div
                  className={`h-full rounded-full ${getHealthColor(value)}`}
                  style={{ width: `${value}%` }}
                />
              </div>
              <span className="text-sm font-medium">{value}%</span>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};