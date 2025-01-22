import React from 'react';
import { Link } from 'react-router-dom';
import { Server, Activity, Shield, Zap, BarChart, Globe, Cpu } from 'lucide-react';
import { motion } from 'framer-motion';
import { DemoChart } from '../components/home/DemoChart';
import { FeatureCard } from '../components/home/FeatureCard';

export const HomePage = () => {
  return (
    <div className="min-h-screen bg-gradient-to-b from-gray-50 to-white">
      {/* Hero Section */}
      <div className="relative overflow-hidden">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-24">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8 }}
            className="text-center"
          >
            <h1 className="text-5xl md:text-6xl font-extrabold text-gray-900 tracking-tight">
              <span className="block">Monitor Your Servers</span>
              <span className="block text-indigo-600 mt-2">With Real-Time Insights</span>
            </h1>
            <p className="mt-6 max-w-lg mx-auto text-xl text-gray-500">
              Get detailed metrics, real-time alerts, and powerful insights for your entire infrastructure.
            </p>
            <div className="mt-10 flex justify-center gap-4">
              <Link
                to="/signup"
                className="px-8 py-3 border border-transparent text-base font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 md:py-4 md:text-lg md:px-10 transition-colors"
              >
                Get Started
              </Link>
              <Link
                to="/login"
                className="px-8 py-3 border border-gray-300 text-base font-medium rounded-md text-indigo-600 bg-white hover:bg-gray-50 md:py-4 md:text-lg md:px-10 transition-colors"
              >
                Live Demo
              </Link>
            </div>
          </motion.div>
        </div>
      </div>

      {/* Demo Charts Section */}
      <div className="py-16 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold text-gray-900">Real-Time Monitoring</h2>
            <p className="mt-4 text-xl text-gray-600">Track your server performance with beautiful visualizations</p>
          </div>
          {/* <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            <DemoChart />
            <DemoChart />
            <DemoChart />
          </div> */}
        </div>
      </div>

      {/* Features Grid */}
      <div className="py-16">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold text-gray-900">Powerful Features</h2>
            <p className="mt-4 text-xl text-gray-600">Everything you need to monitor your infrastructure</p>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            <FeatureCard
              icon={Server}
              title="Server Monitoring"
              description="Track CPU, memory, disk, and network metrics in real-time."
              delay={0.1}
            />
            <FeatureCard
              icon={Globe}
              title="Global Traffic Analysis"
              description="Monitor incoming traffic and requests from around the world."
              delay={0.2}
            />
            <FeatureCard
              icon={Shield}
              title="Security Alerts"
              description="Get instant notifications for security and performance issues."
              delay={0.3}
            />
            <FeatureCard
              icon={BarChart}
              title="Advanced Analytics"
              description="Detailed insights and trends for your infrastructure."
              delay={0.4}
            />
            <FeatureCard
              icon={Cpu}
              title="Resource Optimization"
              description="Identify bottlenecks and optimize server performance."
              delay={0.5}
            />
            <FeatureCard
              icon={Activity}
              title="Health Monitoring"
              description="Continuous monitoring of server health and availability."
              delay={0.6}
            />
          </div>
        </div>
      </div>
    </div>
  );
};