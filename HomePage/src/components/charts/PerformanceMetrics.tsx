import React from 'react';
import { Bar } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);

export function PerformanceMetrics() {
  const data = {
    labels: ['API', 'Database', 'Cache', 'Queue', 'Storage'],
    datasets: [
      {
        label: 'Response Time (ms)',
        data: [45, 82, 15, 32, 68],
        backgroundColor: [
          'rgba(129, 140, 248, 0.8)',
          'rgba(168, 85, 247, 0.8)',
          'rgba(236, 72, 153, 0.8)',
          'rgba(248, 113, 113, 0.8)',
          'rgba(251, 146, 60, 0.8)',
        ],
        borderRadius: 8,
      },
    ],
  };

  const options = {
    responsive: true,
    plugins: {
      legend: {
        display: false,
      },
      tooltip: {
        callbacks: {
          label: function(context: any) {
            return `${context.parsed.y}ms`;
          }
        }
      }
    },
    scales: {
      y: {
        beginAtZero: true,
        grid: {
          color: 'rgba(255, 255, 255, 0.1)',
        },
        ticks: {
          color: 'rgba(255, 255, 255, 0.7)',
        },
      },
      x: {
        grid: {
          display: false,
        },
        ticks: {
          color: 'rgba(255, 255, 255, 0.7)',
        },
      },
    },
  };

  return (
    <div className="p-4">
      <h3 className="text-lg font-semibold mb-4 text-white">Service Performance</h3>
      <Bar data={data} options={options} />
    </div>
  );
}