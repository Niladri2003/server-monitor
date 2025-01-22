import React from 'react';
import { ResponsiveBar } from '@nivo/bar';
import { format } from 'date-fns';

interface MemoryBarChartProps {
  data: Array<{
    timestamp: number;
    used: number;
    cached: number;
    buffers: number;
    free: number;
  }>;
}

export const MemoryBarChart: React.FC<MemoryBarChartProps> = ({ data }) => {
  const chartData = data.map((d) => ({
    time: format(d.timestamp, 'HH:mm:ss'),
    Used: d.used / 1024,
    Cached: d.cached / 1024,
    Buffers: d.buffers / 1024,
    Free: d.free / 1024,
  }));

  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <h3 className="text-lg font-semibold mb-4">Memory Usage</h3>
      <div className="h-[300px]">
        <ResponsiveBar
          data={chartData}
          keys={['Used', 'Cached', 'Buffers', 'Free']}
          indexBy="time"
          margin={{ top: 20, right: 130, bottom: 50, left: 60 }}
          padding={0.3}
          valueScale={{ type: 'linear' }}
          colors={['#ef4444', '#3b82f6', '#10b981', '#6366f1']}
          borderRadius={4}
          axisTop={null}
          axisRight={null}
          axisBottom={{
            tickSize: 5,
            tickPadding: 5,
            tickRotation: -45,
            legend: 'Time',
            legendPosition: 'middle',
            legendOffset: 40,
          }}
          axisLeft={{
            tickSize: 5,
            tickPadding: 5,
            tickRotation: 0,
            legend: 'Memory (GB)',
            legendPosition: 'middle',
            legendOffset: -50,
          }}
          labelSkipWidth={12}
          labelSkipHeight={12}
          legends={[
            {
              dataFrom: 'keys',
              anchor: 'bottom-right',
              direction: 'column',
              justify: false,
              translateX: 120,
              translateY: 0,
              itemsSpacing: 2,
              itemWidth: 100,
              itemHeight: 20,
              itemDirection: 'left-to-right',
              itemOpacity: 0.85,
              symbolSize: 20,
            },
          ]}
        />
      </div>
    </div>
  );
};