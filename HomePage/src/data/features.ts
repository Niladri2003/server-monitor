import { Cpu, CircuitBoard, Network, HardDrive, Globe, Activity, BarChart, Gauge } from 'lucide-react';
import type { LucideIcon } from 'lucide-react';

export interface Feature {
  icon: LucideIcon;
  title: string;
  description: string;
}

export const features: Feature[] = [
  { icon: Cpu, title: 'CPU Monitoring', description: 'Real-time CPU usage tracking and analysis' },
  { icon: CircuitBoard, title: 'Memory Usage', description: 'Detailed memory allocation and consumption metrics' },
  { icon: Network, title: 'Network Traffic', description: 'Bandwidth monitoring and traffic analysis' },
  { icon: HardDrive, title: 'Disk Usage', description: 'Storage capacity and I/O performance tracking' },
  { icon: Globe, title: 'Global Monitoring', description: 'Worldwide request origin tracking' },
  { icon: Activity, title: 'Performance Metrics', description: 'Comprehensive performance analytics' },
  { icon: BarChart, title: 'Custom Metrics', description: 'Define and track custom server metrics' },
  { icon: Gauge, title: 'Resource Usage', description: 'Detailed resource utilization insights' },
];