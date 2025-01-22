import React from 'react';
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';
import { format } from 'date-fns';

interface MemoryChartProps {
  data: Array<{
    timestamp: number;
    used: number;
    cached: number;
    buffers: number;
    free: number;
  }>;
}

export const MemoryChart: React.FC<MemoryChartProps> = ({ data }) => {
  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <h3 className="text-lg font-semibold mb-4">Memory Usage</h3>
      <div className="h-[300px]">
        <ResponsiveContainer width="100%" height="100%">
          <BarChart data={data} margin={{ top: 5, right: 30, left: 20, bottom: 5 }}>
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis
              dataKey="timestamp"
              tickFormatter={(timestamp) => format(timestamp, 'HH:mm')}
            />
            <YAxis unit="GB" />
            <Tooltip
              labelFormatter={(timestamp) => format(timestamp, 'HH:mm:ss')}
              formatter={(value: number) => [`${(value / 1024).toFixed(2)} GB`]}
            />
            <Bar dataKey="used" stackId="a" fill="#ef4444" name="Used" />
            <Bar dataKey="cached" stackId="a" fill="#3b82f6" name="Cached" />
            <Bar dataKey="buffers" stackId="a" fill="#10b981" name="Buffers" />
            <Bar dataKey="free" stackId="a" fill="#6366f1" name="Free" />
          </BarChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
};