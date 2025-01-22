import React from 'react';
import { NavLink, useLocation } from 'react-router-dom';
import { ChevronDown, ChevronRight } from 'lucide-react';

interface DocSection {
  title: string;
  path: string;
  icon?: React.ReactNode;
  children?: DocSection[];
}

const docSections: DocSection[] = [
  {
    title: 'Getting Started',
    path: '/docs/getting-started',
    children: [
      { title: 'Installation', path: '/docs/getting-started/installation' },
      { title: 'Configuration', path: '/docs/getting-started/configuration' },
      { title: 'Quick Start', path: '/docs/getting-started/quick-start' },
    ],
  },
  {
    title: 'Core Concepts',
    path: '/docs/core-concepts',
    children: [
      { title: 'Architecture', path: '/docs/core-concepts/architecture' },
      { title: 'Monitoring', path: '/docs/core-concepts/monitoring' },
      { title: 'Alerting', path: '/docs/core-concepts/alerting' },
    ],
  },
  {
    title: 'API Reference',
    path: '/docs/api',
    children: [
      { title: 'Authentication', path: '/docs/api/authentication' },
      { title: 'Endpoints', path: '/docs/api/endpoints' },
      { title: 'Rate Limits', path: '/docs/api/rate-limits' },
    ],
  },
  {
    title: 'Guides',
    path: '/docs/guides',
    children: [
      { title: 'Best Practices', path: '/docs/guides/best-practices' },
      { title: 'Troubleshooting', path: '/docs/guides/troubleshooting' },
      { title: 'Security', path: '/docs/guides/security' },
    ],
  },
];

const NavItem: React.FC<{ section: DocSection; level?: number }> = ({ 
  section, 
  level = 0 
}) => {
  const [isOpen, setIsOpen] = React.useState(true);
  const location = useLocation();
  const isActive = location.pathname.startsWith(section.path);
  const hasChildren = section.children && section.children.length > 0;

  return (
    <div className="space-y-1">
      <NavLink
        to={section.path}
        className={({ isActive }) =>
          `flex items-center justify-between px-4 py-2 text-sm font-medium rounded-md ${
            isActive
              ? 'text-indigo-600 bg-indigo-50'
              : 'text-gray-600 hover:text-indigo-600 hover:bg-gray-50'
          }`
        }
        style={{ paddingLeft: `${level * 1}rem` }}
      >
        <span>{section.title}</span>
        {hasChildren && (
          <button
            onClick={(e) => {
              e.preventDefault();
              setIsOpen(!isOpen);
            }}
            className="p-1"
          >
            {isOpen ? (
              <ChevronDown className="w-4 h-4" />
            ) : (
              <ChevronRight className="w-4 h-4" />
            )}
          </button>
        )}
      </NavLink>
      {hasChildren && isOpen && (
        <div className="ml-4">
          {section.children.map((child) => (
            <NavItem key={child.path} section={child} level={level + 1} />
          ))}
        </div>
      )}
    </div>
  );
};

export const DocsNavigation = () => {
  return (
    <nav className="space-y-1">
      {docSections.map((section) => (
        <NavItem key={section.path} section={section} />
      ))}
    </nav>
  );
};