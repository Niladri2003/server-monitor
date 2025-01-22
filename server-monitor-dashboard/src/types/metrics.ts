export interface ServerMetrics {
  id: string;
  hostname: string;
  timestamp: number;
  cpuCores: CpuCoreMetric[];
  totalCpuUsage: number;
  memoryUsage: MemoryMetric;
  networkUsage: NetworkMetric;
}

export interface CpuCoreMetric {
  coreId: number;
  usage: number;
}

export interface MemoryMetric {
  total: number;
  used: number;
  free: number;
  cached: number;
  buffers: number;
}

export interface NetworkMetric {
  bytesReceived: number;
  bytesSent: number;
  packetsReceived: number;
  packetsSent: number;
}