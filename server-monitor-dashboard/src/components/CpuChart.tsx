import React from 'react';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';
import { format } from 'date-fns';

interface CpuChartProps {
  data: Array<{
    timestamp: number;
    usage: number;
  }>;
}

export const CpuChart: React.FC<CpuChartProps> = ({ data }) => {
  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <h3 className="text-lg font-semibold mb-4">CPU Utilization</h3>
      <div className="h-[300px]">
        <ResponsiveContainer width="100%" height="100%">
          <LineChart data={data} margin={{ top: 5, right: 30, left: 20, bottom: 5 }}>
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis
              dataKey="timestamp"
              tickFormatter={(timestamp) => format(timestamp, 'HH:mm')}
            />
            <YAxis unit="%" domain={[0, 100]} />
            <Tooltip
              labelFormatter={(timestamp) => format(timestamp, 'HH:mm:ss')}
              formatter={(value: number) => [`${value.toFixed(1)}%`, 'Usage']}
            />
            <Line
              type="monotone"
              dataKey="usage"
              stroke="#6366f1"
              strokeWidth={2}
              dot={false}
            />
          </LineChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
};