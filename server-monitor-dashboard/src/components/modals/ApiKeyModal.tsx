import React, { useState } from 'react';
import { Modal } from './Modal';
import { Copy } from 'lucide-react';

interface ApiKeyModalProps {
  isOpen: boolean;
  onClose: () => void;
}

export const ApiKeyModal: React.FC<ApiKeyModalProps> = ({ isOpen, onClose }) => {
  const [name, setName] = useState('');
  const [generatedKey, setGeneratedKey] = useState('');

  const handleGenerate = () => {
    const key = `sk_${Math.random().toString(36).substr(2, 9)}`;
    setGeneratedKey(key);
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Generate API Key">
      <div className="space-y-4">
        <div>
          <label className="block text-sm font-medium text-gray-700">Key Name</label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
            placeholder="Production Server 1"
          />
        </div>
        {!generatedKey ? (
          <button
            onClick={handleGenerate}
            className="w-full bg-indigo-600 text-white px-4 py-2 rounded-md hover:bg-indigo-700"
          >
            Generate Key
          </button>
        ) : (
          <div className="space-y-2">
            <label className="block text-sm font-medium text-gray-700">API Key</label>
            <div className="flex items-center space-x-2">
              <code className="flex-1 p-2 bg-gray-100 rounded">{generatedKey}</code>
              <button
                onClick={() => navigator.clipboard.writeText(generatedKey)}
                className="p-2 text-gray-600 hover:text-indigo-600"
              >
                <Copy className="w-5 h-5" />
              </button>
            </div>
          </div>
        )}
      </div>
    </Modal>
  );
};