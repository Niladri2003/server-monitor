import React from 'react';

interface StatsCardProps {
  title: string;
  value: string | number;
  trend?: {
    value: string;
    direction: 'up' | 'down';
  };
  subtitle?: string;
}

export const StatsCard: React.FC<StatsCardProps> = ({ title, value, trend, subtitle }) => {
  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <h3 className="text-sm font-medium text-gray-500">{title}</h3>
      <p className="text-2xl font-semibold mt-2">{value}</p>
      {trend && (
        <div className={`text-sm ${trend.direction === 'up' ? 'text-green-600' : 'text-red-600'} mt-2`}>
          {trend.direction === 'up' ? '↑' : '↓'} {trend.value}
        </div>
      )}
      {subtitle && <div className="text-sm text-gray-600 mt-2">{subtitle}</div>}
    </div>
  );
};