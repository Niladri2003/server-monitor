import React from 'react';
import { Settings, Bell, Shield, Clock } from 'lucide-react';

export const SettingsPage = () => {
  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold text-gray-900">Settings</h1>

      <div className="bg-white rounded-lg shadow divide-y">
        {/* General Settings */}
        <div className="p-6">
          <div className="flex items-center space-x-2 mb-4">
            <Settings className="w-5 h-5 text-indigo-600" />
            <h2 className="text-lg font-semibold">General Settings</h2>
          </div>
          
          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="font-medium">Time Zone</p>
                <p className="text-sm text-gray-500">Set your preferred time zone for reports</p>
              </div>
              <select className="rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
                <option>UTC</option>
                <option>America/New_York</option>
                <option>Europe/London</option>
              </select>
            </div>
          </div>
        </div>

        {/* Notifications */}
        <div className="p-6">
          <div className="flex items-center space-x-2 mb-4">
            <Bell className="w-5 h-5 text-indigo-600" />
            <h2 className="text-lg font-semibold">Notifications</h2>
          </div>
          
          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="font-medium">Email Notifications</p>
                <p className="text-sm text-gray-500">Receive alerts via email</p>
              </div>
              <label className="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" className="sr-only peer" checked />
                <div className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-indigo-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-indigo-600"></div>
              </label>
            </div>
          </div>
        </div>

        {/* Security */}
        <div className="p-6">
          <div className="flex items-center space-x-2 mb-4">
            <Shield className="w-5 h-5 text-indigo-600" />
            <h2 className="text-lg font-semibold">Security</h2>
          </div>
          
          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="font-medium">Two-Factor Authentication</p>
                <p className="text-sm text-gray-500">Add an extra layer of security</p>
              </div>
              <button className="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700">
                Enable 2FA
              </button>
            </div>
          </div>
        </div>

        {/* Data Retention */}
        <div className="p-6">
          <div className="flex items-center space-x-2 mb-4">
            <Clock className="w-5 h-5 text-indigo-600" />
            <h2 className="text-lg font-semibold">Data Retention</h2>
          </div>
          
          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="font-medium">Metrics Retention Period</p>
                <p className="text-sm text-gray-500">How long to keep detailed metrics</p>
              </div>
              <select className="rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
                <option>30 days</option>
                <option>60 days</option>
                <option>90 days</option>
                <option>180 days</option>
              </select>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};