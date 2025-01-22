import React from 'react';
import { PlusCircle } from 'lucide-react';

interface ChartOption {
  id: string;
  name: string;
  description: string;
}

interface ChartPickerProps {
  onAddChart: (chartType: string) => void;
}

const chartOptions: ChartOption[] = [
  { id: 'cpu', name: 'CPU Usage', description: 'Monitor CPU utilization over time' },
  { id: 'memory', name: 'Memory Usage', description: 'Track memory consumption' },
  { id: 'disk', name: 'Disk Usage', description: 'View disk space allocation' },
  { id: 'process', name: 'Process Monitor', description: 'Monitor top processes' },
  { id: 'network', name: 'Network Traffic', description: 'Track network I/O' },
];

export const ChartPicker: React.FC<ChartPickerProps> = ({ onAddChart }) => {
  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <h3 className="text-lg font-semibold mb-4">Add Charts</h3>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {chartOptions.map((chart) => (
          <button
            key={chart.id}
            onClick={() => onAddChart(chart.id)}
            className="flex items-center p-4 border rounded-lg hover:border-indigo-500 hover:bg-indigo-50 transition-colors"
          >
            <PlusCircle className="w-5 h-5 text-indigo-600 mr-3" />
            <div className="text-left">
              <h4 className="font-medium">{chart.name}</h4>
              <p className="text-sm text-gray-500">{chart.description}</p>
            </div>
          </button>
        ))}
      </div>
    </div>
  );
};