import React from 'react';
import { TrendingUp } from 'lucide-react';
import { ResponsiveLine } from '@nivo/line';

interface PredictionData {
  metric: string;
  current: number;
  predicted: number;
  timeline: Array<{
    timestamp: string;
    value: number;
    predicted: boolean;
  }>;
}

interface ResourcePredictionProps {
  predictions: PredictionData[];
}

export const ResourcePrediction: React.FC<ResourcePredictionProps> = ({ predictions }) => {
  const formatPredictionData = (data: PredictionData) => {
    return [
      {
        id: 'actual',
        data: data.timeline
          .filter((point) => !point.predicted)
          .map((point) => ({
            x: point.timestamp,
            y: point.value,
          })),
      },
      {
        id: 'predicted',
        data: data.timeline
          .filter((point) => point.predicted)
          .map((point) => ({
            x: point.timestamp,
            y: point.value,
          })),
      },
    ];
  };

  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <div className="flex items-center justify-between mb-4">
        <h3 className="text-lg font-semibold">Resource Predictions</h3>
        <TrendingUp className="w-5 h-5 text-indigo-600" />
      </div>
      <div className="space-y-6">
        {predictions.map((prediction) => (
          <div key={prediction.metric} className="space-y-2">
            <div className="flex items-center justify-between">
              <span className="text-sm font-medium capitalize">
                {prediction.metric}
              </span>
              <span className="text-sm text-gray-500">
                Predicted: {prediction.predicted}%
              </span>
            </div>
            <div className="h-32">
              <ResponsiveLine
                data={formatPredictionData(prediction)}
                margin={{ top: 10, right: 10, bottom: 20, left: 40 }}
                xScale={{ type: 'point' }}
                yScale={{ type: 'linear', min: 0, max: 100 }}
                curve="monotoneX"
                enablePoints={false}
                enableGridX={false}
                colors={['#6366f1', '#e5e7eb']}
                lineWidth={2}
                enableArea={true}
                areaOpacity={0.1}
                axisBottom={null}
                axisLeft={{
                  tickSize: 0,
                  tickPadding: 5,
                  tickRotation: 0,
                  format: (value) => `${value}%`,
                }}
              />
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};