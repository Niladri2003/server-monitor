import React from 'react';
import { Server } from 'lucide-react';

interface ServerListProps {
  servers: Array<{
    id: string;
    hostname: string;
    status: 'online' | 'offline';
  }>;
  selectedServer: string;
  onServerSelect: (serverId: string) => void;
}

export const ServerList: React.FC<ServerListProps> = ({
  servers,
  selectedServer,
  onServerSelect,
}) => {
  return (
    <div className="bg-white rounded-lg shadow-md p-4">
      <h3 className="text-lg font-semibold mb-4">Servers</h3>
      <div className="space-y-2">
        {servers.map((server) => (
          <button
            key={server.id}
            onClick={() => onServerSelect(server.id)}
            className={`w-full flex items-center p-3 rounded-lg transition-colors ${
              selectedServer === server.id
                ? 'bg-indigo-50 text-indigo-600'
                : 'hover:bg-gray-50'
            }`}
          >
            <Server className="w-5 h-5 mr-3" />
            <div className="flex-1 text-left">
              <p className="font-medium">{server.hostname}</p>
              <p className="text-sm text-gray-500">ID: {server.id}</p>
            </div>
            <div
              className={`w-2 h-2 rounded-full ${
                server.status === 'online' ? 'bg-green-500' : 'bg-red-500'
              }`}
            />
          </button>
        ))}
      </div>
    </div>
  );
};