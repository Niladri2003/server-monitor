import React from 'react';
import { NavLink } from 'react-router-dom';
import { 
  LayoutDashboard, 
  Server, 
  Key, 
  Settings, 
  Bell,
  Users,
  Network
} from 'lucide-react';
import { ServerList } from '../ServerList';
import { useLocation } from 'react-router-dom';

interface SidebarProps {
  selectedServer: string;
  onServerSelect: (serverId: string) => void;
}

export const Sidebar: React.FC<SidebarProps> = ({
  selectedServer,
  onServerSelect,
}) => {
  const location = useLocation();
  const navigationItems = [
    { path: '/dashboard', label: 'Dashboard', icon: LayoutDashboard },
    { path: '/servers', label: 'Servers', icon: Server },
    { path: '/port-monitoring', label: 'Port Monitoring', icon: Network },
    { path: '/api-keys', label: 'API Keys', icon: Key },
    { path: '/alerts', label: 'Alerts', icon: Bell },
    { path: '/teams', label: 'Team Management', icon: Users },
    { path: '/settings', label: 'Settings', icon: Settings },
  ];

  return (
    <div className="h-full w-64 bg-white shadow-md p-4 space-y-6">
      {/* Logo/Brand */}
      <div className="flex items-center px-3 mb-6">
        <Server className="w-6 h-6 text-indigo-600 mr-2" />
        <span className="text-xl font-bold text-gray-900">Sysmos</span>
      </div>

      {/* Navigation */}
      <div className="space-y-1">
        {navigationItems.map((item) => {
          const Icon = item.icon;
          return (
            <NavLink
              key={item.path}
              to={item.path}
              className={({ isActive }) =>
                `w-full flex items-center p-3 rounded-lg transition-colors ${
                  isActive
                    ? 'bg-indigo-50 text-indigo-600'
                    : 'text-gray-600 hover:bg-gray-50'
                }`
              }
            >
              <Icon className="w-5 h-5 mr-3" />
              <span className="font-medium">{item.label}</span>
            </NavLink>
          );
        })}
      </div>

      {/* Server List - Only show on dashboard */}
      {location.pathname === '/dashboard' && (
        <div className="mt-6 pt-6 border-t border-gray-200">
          <ServerList
            servers={[
              { id: 'srv-1', hostname: 'prod-app-1', status: 'online' },
              { id: 'srv-2', hostname: 'prod-app-2', status: 'online' },
              { id: 'srv-3', hostname: 'prod-db-1', status: 'offline' },
            ]}
            selectedServer={selectedServer}
            onServerSelect={onServerSelect}
          />
        </div>
      )}
    </div>
  );
};