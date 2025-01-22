import React from 'react';
import { ResponsiveLine } from '@nivo/line';
import { motion } from 'framer-motion';

const generateDemoData = () => {
  return Array.from({ length: 24 }, (_, i) => ({
    x: `${i}:00`,
    y: Math.floor(Math.random() * 60) + 20,
  }));
};

export const DemoChart = () => {
  const data = [{
    id: 'server metrics',
    data: generateDemoData()
  }];

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      whileInView={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.8 }}
      viewport={{ once: true }}
      className="bg-white p-6 rounded-xl shadow-lg"
    >
      <h3 className="text-lg font-semibold mb-4">Live Server Metrics</h3>
      <div className="h-[200px]">
        <ResponsiveLine
          data={data}
          margin={{ top: 20, right: 20, bottom: 40, left: 40 }}
          xScale={{ type: 'point' }}
          yScale={{ type: 'linear', min: 'auto', max: 'auto' }}
          curve="monotoneX"
          axisTop={null}
          axisRight={null}
          axisBottom={{
            tickSize: 5,
            tickPadding: 5,
            tickRotation: -45,
          }}
          axisLeft={{
            tickSize: 5,
            tickPadding: 5,
            tickRotation: 0,
          }}
          enablePoints={false}
          enableArea={true}
          areaOpacity={0.15}
          colors={['#6366f1']}
          theme={{
            axis: {
              ticks: {
                text: {
                  fontSize: 12,
                },
              },
            },
            grid: {
              line: {
                stroke: '#e2e8f0',
                strokeWidth: 1,
              },
            },
          }}
        />
      </div>
    </motion.div>
  );
};