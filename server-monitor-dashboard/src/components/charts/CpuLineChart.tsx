import React from 'react';
import { ResponsiveLine } from '@nivo/line';
import { format } from 'date-fns';

interface CpuLineChartProps {
  data: Array<{
    timestamp: number;
    usage: number;
  }>;
}

export const CpuLineChart: React.FC<CpuLineChartProps> = ({ data }) => {
  const chartData = [
    {
      id: 'cpu',
      data: data.map((d) => ({
        x: format(d.timestamp, 'HH:mm:ss'),
        y: d.usage,
      })),
    },
  ];

  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <h3 className="text-lg font-semibold mb-4">CPU Utilization</h3>
      <div className="h-[300px]">
        <ResponsiveLine
          data={chartData}
          margin={{ top: 20, right: 20, bottom: 50, left: 50 }}
          xScale={{ type: 'point' }}
          yScale={{ type: 'linear', min: 0, max: 100 }}
          curve="monotoneX"
          axisTop={null}
          axisRight={null}
          axisBottom={{
            tickSize: 5,
            tickPadding: 5,
            tickRotation: -45,
            legend: 'Time',
            legendOffset: 40,
          }}
          axisLeft={{
            tickSize: 5,
            tickPadding: 5,
            tickRotation: 0,
            legend: 'Usage (%)',
            legendOffset: -40,
          }}
          enablePoints={false}
          enableArea={true}
          areaOpacity={0.15}
          colors={['#6366f1']}
          theme={{
            axis: {
              ticks: {
                text: {
                  fontSize: 12,
                },
              },
            },
            grid: {
              line: {
                stroke: '#e2e8f0',
                strokeWidth: 1,
              },
            },
          }}
        />
      </div>
    </div>
  );
};