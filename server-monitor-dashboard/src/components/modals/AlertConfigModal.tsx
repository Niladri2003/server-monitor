import React, { useState } from 'react';
import { Modal } from './Modal';

interface AlertConfigModalProps {
  isOpen: boolean;
  onClose: () => void;
}

export const AlertConfigModal: React.FC<AlertConfigModalProps> = ({ isOpen, onClose }) => {
  const [threshold, setThreshold] = useState('90');
  const [metric, setMetric] = useState('cpu');

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Configure Alerts">
      <div className="space-y-4">
        <div>
          <label className="block text-sm font-medium text-gray-700">Metric</label>
          <select
            value={metric}
            onChange={(e) => setMetric(e.target.value)}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
          >
            <option value="cpu">CPU Usage</option>
            <option value="memory">Memory Usage</option>
            <option value="disk">Disk Usage</option>
            <option value="network">Network Usage</option>
          </select>
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700">Threshold (%)</label>
          <input
            type="number"
            value={threshold}
            onChange={(e) => setThreshold(e.target.value)}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
          />
        </div>
        <button
          onClick={onClose}
          className="w-full bg-indigo-600 text-white px-4 py-2 rounded-md hover:bg-indigo-700"
        >
          Save Configuration
        </button>
      </div>
    </Modal>
  );
};