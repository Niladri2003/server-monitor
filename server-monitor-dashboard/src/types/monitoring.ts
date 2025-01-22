export interface RequestLocation {
  name: string;
  coordinates: [number, number]; // [longitude, latitude]
  requests: number;
}

export interface MonitoringStats {
  totalRequests: number;
  activePorts: number;
  avgResponseTime: number;
  failedRequests: number;
}