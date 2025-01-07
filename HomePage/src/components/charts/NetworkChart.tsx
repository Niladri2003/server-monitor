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

export function NetworkChart() {
  const data = {
    labels: ['US', 'EU', 'Asia', 'AU', 'SA'],
    datasets: [
      {
        label: 'Incoming Traffic (Mb/s)',
        data: Array.from({ length: 5 }, () => Math.floor(Math.random() * 1000)),
        backgroundColor: 'rgba(99, 102, 241, 0.8)',
      },
      {
        label: 'Outgoing Traffic (Mb/s)',
        data: Array.from({ length: 5 }, () => Math.floor(Math.random() * 1000)),
        backgroundColor: 'rgba(45, 212, 191, 0.8)',
      },
    ],
  };

  const options = {
    responsive: true,
    plugins: {
      legend: {
        position: 'bottom' as const,
      },
    },
    scales: {
      y: {
        beginAtZero: true,
        grid: {
          display: false,
        },
      },
      x: {
        grid: {
          display: false,
        },
      },
    },
    animation: {
      duration: 2000,
    },
  };

  return <Bar data={data} options={options} />;
}