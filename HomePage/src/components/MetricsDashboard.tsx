import React from 'react';
import { DashboardCard } from './dashboard/DashboardCard';
import { ServerStatus } from './dashboard/ServerStatus';
import { AlertsTimeline } from './dashboard/AlertsTimeline';
import { ResourceMetrics } from './dashboard/ResourceMetrics';
import { ContactButton } from './contact/ContactButton';
import { Monitor, Activity } from 'lucide-react';

export default function MetricsDashboard() {
  return (
    <section id="metrics" className="py-20 bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center mb-16">
          <h2 className="text-4xl font-bold text-gray-900">Interactive Analytics Dashboard</h2>
          <p className="mt-4 text-xl text-gray-600">Real-time infrastructure insights at your fingertips</p>
        </div>
        
        <div className="relative">
          {/* Monitor Frame */}
          <div className="relative mx-auto max-w-6xl">
            <div className="bg-gray-900 rounded-t-3xl p-8 pb-4">
              <div className="flex items-center justify-between mb-4">
                <div className="flex items-center space-x-2">
                  <Activity className="h-5 w-5 text-green-400" />
                  <span className="text-white font-medium">System Monitor</span>
                </div>
                <div className="flex space-x-2">
                  <div className="w-3 h-3 rounded-full bg-red-500"></div>
                  <div className="w-3 h-3 rounded-full bg-yellow-500"></div>
                  <div className="w-3 h-3 rounded-full bg-green-500"></div>
                </div>
              </div>
            </div>
            
            <div className="bg-white rounded-b-3xl shadow-2xl p-8">
              <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
                <div className="lg:col-span-2">
                  <ResourceMetrics />
                </div>
                <div className="space-y-8">
                  <ServerStatus />
                  <AlertsTimeline />
                </div>
              </div>
            </div>
          </div>

          {/* Decorative Elements */}
          <div className="absolute -top-10 -left-10 w-72 h-72 bg-indigo-500 rounded-full mix-blend-multiply filter blur-xl opacity-10 animate-blob"></div>
          <div className="absolute -bottom-10 -right-10 w-72 h-72 bg-purple-500 rounded-full mix-blend-multiply filter blur-xl opacity-10 animate-blob"></div>
        </div>

        <div className="mt-16 text-center">
          <ContactButton />
        </div>
      </div>
    </section>
  );
}