import React from 'react';
import { SystemMetrics } from './charts/SystemMetrics';
import { PerformanceMetrics } from './charts/PerformanceMetrics';

export function MetricsPreview() {
  return (
    <div className="relative mt-16">
      <div className="bg-white/10 backdrop-blur-lg rounded-2xl p-8 shadow-2xl">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          <SystemMetrics />
          <PerformanceMetrics />
        </div>
      </div>
      {/* Decorative elements */}
      <div className="absolute -top-4 -left-4 w-72 h-72 bg-indigo-500 rounded-full mix-blend-multiply filter blur-xl opacity-20 animate-blob"></div>
      <div className="absolute -bottom-8 -right-4 w-72 h-72 bg-purple-500 rounded-full mix-blend-multiply filter blur-xl opacity-20 animate-blob animation-delay-2000"></div>
    </div>
  );
}