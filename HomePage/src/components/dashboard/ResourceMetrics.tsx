import React from 'react';
import { Line } from 'react-chartjs-2';
import { DashboardCard } from './DashboardCard';
import { Chart as ChartJS, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend, Filler } from 'chart.js';

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend, Filler);

export function ResourceMetrics() {
  const data = {
    labels: Array.from({ length: 12 }, (_, i) => `${i * 2}:00`),
    datasets: [
      {
        label: 'CPU Usage',
        data: Array.from({ length: 12 }, () => Math.floor(Math.random() * 40) + 30),
        borderColor: 'rgba(99, 102, 241, 1)',
        backgroundColor: 'rgba(99, 102, 241, 0.1)',
        fill: true,
        tension: 0.4,
      },
      {
        label: 'Memory Usage',
        data: Array.from({ length: 12 }, () => Math.floor(Math.random() * 30) + 50),
        borderColor: 'rgba(168, 85, 247, 1)',
        backgroundColor: 'rgba(168, 85, 247, 0.1)',
        fill: true,
        tension: 0.4,
      },
      {
        label: 'Network I/O',
        data: Array.from({ length: 12 }, () => Math.floor(Math.random() * 50) + 20),
        borderColor: 'rgba(236, 72, 153, 1)',
        backgroundColor: 'rgba(236, 72, 153, 0.1)',
        fill: true,
        tension: 0.4,
      },
    ],
  };

  const options = {
    responsive: true,
    interaction: {
      mode: 'index' as const,
      intersect: false,
    },
    plugins: {
      legend: {
        position: 'top' as const,
      },
      tooltip: {
        animation: {
          duration: 200,
        },
      },
    },
    scales: {
      y: {
        beginAtZero: true,
        grid: {
          color: 'rgba(0, 0, 0, 0.05)',
        },
      },
      x: {
        grid: {
          display: false,
        },
      },
    },
  };

  return (
    <DashboardCard title="Resource Utilization (24h)" className="h-full">
      <Line data={data} options={options} className="h-[400px]" />
    </DashboardCard>
  );
}