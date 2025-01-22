import  { useState } from 'react';
import { Cpu, HardDrive, Network, Activity } from 'lucide-react';
import { DndContext, closestCenter, DragEndEvent } from '@dnd-kit/core';
import { SortableContext, arrayMove } from '@dnd-kit/sortable';
import { MetricCard } from '../components/MetricCard';
import { CpuLineChart } from '../components/charts/CpuLineChart';
import { MemoryBarChart } from '../components/charts/MemoryBarChart';
import { NetworkChart } from '../components/NetworkChart';
import { DiskUsageChart } from '../components/charts/DiskUsageChart';
import { ProcessChart } from '../components/charts/ProcessChart';
import { DraggableChart } from '../components/dashboard/DraggableChart';
import { ChartSelectionModal, ChartConfig } from '../components/dashboard/ChartSelectionModal';
import { SystemHealthScore } from '../components/dashboard/SystemHealthScore';
import { AnomalyDetection } from '../components/dashboard/AnomalyDetection';
import { ResourcePrediction } from '../components/dashboard/ResourcePrediction';

// ... (keep existing mock data generation)

const mockAnomalies = [
  {
    id: '1',
    type: 'critical' as const,
    message: 'Unusual CPU spike detected',
    timestamp: '2 minutes ago',
    metric: 'CPU Usage',
  },
  {
    id: '2',
    type: 'warning' as const,
    message: 'Memory usage trending higher than normal',
    timestamp: '5 minutes ago',
    metric: 'Memory Usage',
  },
];

const mockPredictions = [
  {
    metric: 'cpu',
    current: 45,
    predicted: 65,
    timeline: Array.from({ length: 12 }, (_, i) => ({
      timestamp: `${i}:00`,
      value: 45 + Math.random() * 20,
      predicted: i > 6,
    })),
  },
  {
    metric: 'memory',
    current: 62,
    predicted: 78,
    timeline: Array.from({ length: 12 }, (_, i) => ({
      timestamp: `${i}:00`,
      value: 62 + Math.random() * 16,
      predicted: i > 6,
    })),
  },
];
const generateMockTimeSeriesData = (points: number) => {
  const now = Date.now();
  return Array.from({ length: points }, (_, i) => ({
    timestamp: now - (points - i - 1) * 60000,
    usage: 30 + Math.random() * 40,
    used: Math.random() * 8 * 1024,
    cached: Math.random() * 4 * 1024,
    buffers: Math.random() * 2 * 1024,
    free: Math.random() * 2 * 1024,
    bytesReceived: Math.random() * 1024 * 1024 * 100,
    bytesSent: Math.random() * 1024 * 1024 * 50,
  }));
};

const mockDiskData = [
  { name: 'Used', value: 250, color: '#ef4444' },
  { name: 'Free', value: 750, color: '#22c55e' },
];

const mockProcessData = [
  { name: 'nginx', cpu: 25, memory: 15 },
  { name: 'mongodb', cpu: 40, memory: 30 },
  { name: 'node', cpu: 35, memory: 25 },
  { name: 'redis', cpu: 15, memory: 10 },
];

interface DashboardChart extends ChartConfig {
  id: string;
}

// ... (keep existing DashboardChart interface and other type definitions)

export const DashboardPage = () => {
  // ... (keep existing state and handlers)
  const mockData = generateMockTimeSeriesData(30);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [charts, setCharts] = useState<DashboardChart[]>([
    { id: 'cpu-1', type: 'cpu', title: 'CPU Usage' },
    { id: 'memory-1', type: 'memory', title: 'Memory Usage' },
  ]);

  const handleDragEnd = (event: DragEndEvent) => {
    const { active, over } = event;

    if (over && active.id !== over.id) {
      setCharts((items) => {
        const oldIndex = items.findIndex((i) => i.id === active.id);
        const newIndex = items.findIndex((i) => i.id === over.id);
        return arrayMove(items, oldIndex, newIndex);
      });
    }
  };

  const handleAddChart = (config: ChartConfig) => {
    const newChart: DashboardChart = {
      ...config,
      id: `${config.type}-${charts.length + 1}`,
    };
    setCharts([...charts, newChart]);
  };

  const renderChart = (chart: DashboardChart) => {
    switch (chart.type) {
      case 'cpu':
        return <CpuLineChart data={mockData} />;
      case 'memory':
        return <MemoryBarChart data={mockData} />;
      case 'disk':
        return <DiskUsageChart data={mockDiskData} />;
      case 'process':
        return <ProcessChart data={mockProcessData} />;
      case 'network':
        return <NetworkChart data={mockData} />;
      default:
        return null;
    }
  };
  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-900">Dashboard</h1>
        <button 
          onClick={() => setIsModalOpen(true)}
          className="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700"
        >
          Add Chart
        </button>
      </div>
      
      {/* System Overview Section */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <SystemHealthScore
          score={85}
          metrics={{
            cpu: 45,
            memory: 62,
            disk: 78,
            network: 32,
          }}
        />
        <AnomalyDetection anomalies={mockAnomalies} />
        <ResourcePrediction predictions={mockPredictions} />
      </div>

      {/* Metric Cards */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <MetricCard
          title="CPU Usage"
          value="45%"
          icon={<Cpu className="w-5 h-5" />}
          trend="up"
          trendValue="2.5%"
        />
        <MetricCard
          title="Disk I/O"
          value="120 MB/s"
          icon={<HardDrive className="w-5 h-5" />}
          trend="down"
          trendValue="5%"
        />
        <MetricCard
          title="Network In"
          value="2.4 MB/s"
          icon={<Network className="w-5 h-5" />}
        />
        <MetricCard
          title="Network Out"
          value="1.8 MB/s"
          icon={<Activity className="w-5 h-5" />}
        />
      </div>

      {/* Draggable Charts */}
      <DndContext collisionDetection={closestCenter} onDragEnd={handleDragEnd}>
        <SortableContext items={charts.map(c => c.id)}>
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
            {charts.map((chart) => (
              <DraggableChart key={chart.id} id={chart.id}>
                {renderChart(chart)}
              </DraggableChart>
            ))}
          </div>
        </SortableContext>
      </DndContext>

      {/* Chart Selection Modal */}
      <ChartSelectionModal
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        onAddChart={handleAddChart}
      />
    </div>
  );
};