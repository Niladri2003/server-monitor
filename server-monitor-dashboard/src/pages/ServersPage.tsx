import React from 'react';
import { Server, Circle, ArrowUpDown, Clock, Cpu, MemoryStick } from 'lucide-react';

export const ServersPage = () => {
  const mockServers = [
    {
      id: 'srv-1',
      hostname: 'prod-app-1',
      status: 'online',
      ip: '10.0.1.101',
      uptime: '15d 4h',
      cpu: '45%',
      memory: '62%',
    },
    {
      id: 'srv-2',
      hostname: 'prod-app-2',
      status: 'online',
      ip: '10.0.1.102',
      uptime: '8d 12h',
      cpu: '38%',
      memory: '55%',
    },
    {
      id: 'srv-3',
      hostname: 'prod-db-1',
      status: 'offline',
      ip: '10.0.1.103',
      uptime: '0',
      cpu: '0%',
      memory: '0%',
    },
  ];

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-900">Servers</h1>
        <button className="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700">
          Add New Server
        </button>
      </div>

      <div className="bg-white rounded-lg shadow overflow-hidden">
        <div className="overflow-x-auto">
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Status
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Hostname
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  IP Address
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Uptime
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  CPU
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Memory
                </th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              {mockServers.map((server) => (
                <tr key={server.id} className="hover:bg-gray-50">
                  <td className="px-6 py-4 whitespace-nowrap">
                    <div className="flex items-center">
                      <Circle
                        className={`w-3 h-3 ${
                          server.status === 'online' ? 'text-green-500' : 'text-red-500'
                        }`}
                      />
                    </div>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <div className="flex items-center">
                      <Server className="w-5 h-5 text-gray-500 mr-2" />
                      <span className="font-medium">{server.hostname}</span>
                    </div>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-gray-500">
                    {server.ip}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <div className="flex items-center text-gray-500">
                      <Clock className="w-4 h-4 mr-1" />
                      {server.uptime}
                    </div>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <div className="flex items-center text-gray-500">
                      <Cpu className="w-4 h-4 mr-1" />
                      {server.cpu}
                    </div>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <div className="flex items-center text-gray-500">
                      <MemoryStick className="w-4 h-4 mr-1" />
                      {server.memory}
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};