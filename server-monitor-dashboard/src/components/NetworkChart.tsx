import React from 'react';
import { ResponsiveLine } from '@nivo/line';
import { format } from 'date-fns';

interface NetworkChartProps {
  data: Array<{
    timestamp: number;
    bytesReceived: number;
    bytesSent: number;
  }>;
}

export const NetworkChart: React.FC<NetworkChartProps> = ({ data }) => {
  const formatBytes = (bytes: number) => {
    if (bytes === 0) return '0 B';
    const k = 1024;
    const sizes = ['B', 'KB', 'MB', 'GB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return `${(bytes / Math.pow(k, i)).toFixed(2)} ${sizes[i]}`;
  };

  const chartData = [
    {
      id: 'Received',
      data: data.map((d) => ({
        x: format(d.timestamp, 'HH:mm:ss'),
        y: d.bytesReceived,
      })),
    },
    {
      id: 'Sent',
      data: data.map((d) => ({
        x: format(d.timestamp, 'HH:mm:ss'),
        y: d.bytesSent,
      })),
    },
  ];

  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <h3 className="text-lg font-semibold mb-4">Network Traffic</h3>
      <div className="h-[300px]">
        <ResponsiveLine
          data={chartData}
          margin={{ top: 20, right: 110, bottom: 50, left: 60 }}
          xScale={{ type: 'point' }}
          yScale={{
            type: 'linear',
            min: 'auto',
            max: 'auto',
          }}
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
            legend: 'Traffic',
            legendOffset: -50,
            format: formatBytes,
          }}
          pointSize={0}
          enableArea={true}
          areaOpacity={0.15}
          colors={['#3b82f6', '#10b981']}
          enablePoints={false}
          useMesh={true}
          legends={[
            {
              anchor: 'bottom-right',
              direction: 'column',
              justify: false,
              translateX: 100,
              translateY: 0,
              itemsSpacing: 0,
              itemDirection: 'left-to-right',
              itemWidth: 80,
              itemHeight: 20,
              itemOpacity: 0.75,
              symbolSize: 12,
              symbolShape: 'circle',
            },
          ]}
          tooltip={({ point }) => (
            <div className="bg-white p-2 shadow rounded border border-gray-200">
              <strong>{point.serieId}:</strong> {formatBytes(point.data.y as number)}
            </div>
          )}
        />
      </div>
    </div>
  );
};