import React, { useState } from 'react';
import { Modal } from '../modals/Modal';
import { 
  BarChart as BarChartIcon, 
  LineChart as LineChartIcon, 
  PieChart as PieChartIcon,
  Activity,
  Cpu,
  Server,
  Settings,
  Check
} from 'lucide-react';

interface ChartSelectionModalProps {
  isOpen: boolean;
  onClose: () => void;
  onAddChart: (config: ChartConfig) => void;
}

export interface ChartConfig {
  type: string;
  title: string;
  refreshInterval?: number;
  customOptions?: Record<string, any>;
}

const chartTypes = [
  {
    id: 'cpu',
    name: 'CPU Usage',
    description: 'Real-time CPU utilization chart',
    icon: Cpu,
    preview: <div className="bg-gradient-to-r from-indigo-500 to-indigo-600 h-full rounded-lg opacity-80" />,
  },
  {
    id: 'memory',
    name: 'Memory Usage',
    description: 'Memory consumption over time',
    icon: Server,
    preview: <div className="bg-gradient-to-r from-green-500 to-green-600 h-full rounded-lg opacity-80" />,
  },
  {
    id: 'disk',
    name: 'Disk Usage',
    description: 'Storage space allocation',
    icon: PieChartIcon,
    preview: <div className="bg-gradient-to-r from-red-500 to-red-600 h-full rounded-lg opacity-80" />,
  },
  {
    id: 'process',
    name: 'Process Monitor',
    description: 'Top processes by resource usage',
    icon: BarChartIcon,
    preview: <div className="bg-gradient-to-r from-purple-500 to-purple-600 h-full rounded-lg opacity-80" />,
  },
  {
    id: 'network',
    name: 'Network Traffic',
    description: 'Network I/O monitoring',
    icon: Activity,
    preview: <div className="bg-gradient-to-r from-blue-500 to-blue-600 h-full rounded-lg opacity-80" />,
  },
];

const refreshIntervals = [
  { value: 5000, label: '5 seconds' },
  { value: 10000, label: '10 seconds' },
  { value: 30000, label: '30 seconds' },
  { value: 60000, label: '1 minute' },
];

export const ChartSelectionModal: React.FC<ChartSelectionModalProps> = ({
  isOpen,
  onClose,
  onAddChart,
}) => {
  const [selectedChart, setSelectedChart] = useState<string | null>(null);
  const [chartTitle, setChartTitle] = useState('');
  const [refreshInterval, setRefreshInterval] = useState(10000);
  const [step, setStep] = useState<'select' | 'configure'>('select');

  const handleAddChart = () => {
    if (selectedChart && chartTitle) {
      onAddChart({
        type: selectedChart,
        title: chartTitle,
        refreshInterval,
      });
      onClose();
      // Reset state
      setSelectedChart(null);
      setChartTitle('');
      setRefreshInterval(10000);
      setStep('select');
    }
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Add New Chart">
      {step === 'select' ? (
        <div className="space-y-6">
          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
            {chartTypes.map((chart) => {
              const Icon = chart.icon;
              return (
                <button
                  key={chart.id}
                  onClick={() => {
                    setSelectedChart(chart.id);
                    setChartTitle(chart.name);
                    setStep('configure');
                  }}
                  className={`relative p-4 border-2 rounded-lg transition-all ${
                    selectedChart === chart.id
                      ? 'border-indigo-500 bg-indigo-50'
                      : 'border-gray-200 hover:border-indigo-300'
                  }`}
                >
                  <div className="flex items-start space-x-3">
                    <div className="flex-shrink-0">
                      <Icon className="w-6 h-6 text-indigo-600" />
                    </div>
                    <div className="flex-1 text-left">
                      <h3 className="font-medium">{chart.name}</h3>
                      <p className="text-sm text-gray-500">{chart.description}</p>
                    </div>
                  </div>
                  <div className="mt-3 h-24 rounded-lg overflow-hidden">
                    {chart.preview}
                  </div>
                </button>
              );
            })}
          </div>
        </div>
      ) : (
        <div className="space-y-6">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Chart Title
            </label>
            <input
              type="text"
              value={chartTitle}
              onChange={(e) => setChartTitle(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
              placeholder="Enter chart title"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Refresh Interval
            </label>
            <select
              value={refreshInterval}
              onChange={(e) => setRefreshInterval(Number(e.target.value))}
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
            >
              {refreshIntervals.map((interval) => (
                <option key={interval.value} value={interval.value}>
                  {interval.label}
                </option>
              ))}
            </select>
          </div>

          <div className="flex justify-between pt-4">
            <button
              onClick={() => setStep('select')}
              className="px-4 py-2 text-sm font-medium text-gray-700 hover:text-gray-900"
            >
              Back
            </button>
            <button
              onClick={handleAddChart}
              disabled={!chartTitle}
              className="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              Add Chart
            </button>
          </div>
        </div>
      )}
    </Modal>
  );
};