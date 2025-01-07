import React from 'react';
import { Doughnut } from 'react-chartjs-2';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';

ChartJS.register(ArcElement, Tooltip, Legend);

export function MemoryChart() {
  const data = {
    labels: ['Used', 'Free', 'Cached'],
    datasets: [
      {
        data: [65, 20, 15],
        backgroundColor: [
          'rgba(99, 102, 241, 0.8)',
          'rgba(45, 212, 191, 0.8)',
          'rgba(251, 146, 60, 0.8)',
        ],
        borderColor: [
          'rgba(99, 102, 241, 1)',
          'rgba(45, 212, 191, 1)',
          'rgba(251, 146, 60, 1)',
        ],
        borderWidth: 1,
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
    animation: {
      animateRotate: true,
      animateScale: true,
      duration: 2000,
    },
  };

  return <Doughnut data={data} options={options} />;
}