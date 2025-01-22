import React from 'react';
import { Key, Copy, Trash } from 'lucide-react';

export const ApiKeysPage = () => {
  const mockApiKeys = [
    { id: 1, name: 'Production Server 1', key: 'sk_prod_123...abc', created: '2024-03-15' },
    { id: 2, name: 'Staging Server', key: 'sk_stag_456...xyz', created: '2024-03-14' },
  ];

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-900">API Keys</h1>
        <button className="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700">
          Generate New Key
        </button>
      </div>

      <div className="bg-white rounded-lg shadow">
        <div className="p-6">
          <div className="flex items-center space-x-2 mb-4">
            <Key className="w-5 h-5 text-indigo-600" />
            <h2 className="text-lg font-semibold">Your API Keys</h2>
          </div>

          <div className="space-y-4">
            {mockApiKeys.map((apiKey) => (
              <div
                key={apiKey.id}
                className="border rounded-lg p-4 flex items-center justify-between"
              >
                <div>
                  <h3 className="font-medium">{apiKey.name}</h3>
                  <p className="text-sm text-gray-500">Created on {apiKey.created}</p>
                  <code className="text-sm bg-gray-100 px-2 py-1 rounded mt-1">
                    {apiKey.key}
                  </code>
                </div>
                <div className="flex space-x-2">
                  <button className="p-2 text-gray-600 hover:text-indigo-600">
                    <Copy className="w-5 h-5" />
                  </button>
                  <button className="p-2 text-gray-600 hover:text-red-600">
                    <Trash className="w-5 h-5" />
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};