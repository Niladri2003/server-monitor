import React, { useState } from 'react';
import { useLocation } from 'react-router-dom';
import { Search, Menu, X } from 'lucide-react';
import { DocsNavigation } from './DocsNavigation';
import { DocContent } from './DocContent';
import { TableOfContents } from './TableOfContents';

export const DocLayout = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [searchQuery, setSearchQuery] = useState('');
  const location = useLocation();

  return (
    <div className="flex min-h-screen bg-white">
      {/* Mobile menu button */}
      <button
        onClick={() => setIsMenuOpen(!isMenuOpen)}
        className="lg:hidden fixed top-4 left-4 z-50 p-2 bg-white rounded-md shadow-md"
      >
        {isMenuOpen ? <X className="w-6 h-6" /> : <Menu className="w-6 h-6" />}
      </button>

      {/* Left Sidebar */}
      <div
        className={`fixed lg:sticky top-0 h-screen w-64 transform ${
          isMenuOpen ? 'translate-x-0' : '-translate-x-full'
        } lg:translate-x-0 transition-transform duration-300 ease-in-out bg-white border-r border-gray-200`}
      >
        <div className="h-full p-4">
          {/* Search */}
          <div className="mb-6">
            <div className="relative">
              <input
                type="text"
                placeholder="Search documentation..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                className="w-full px-4 py-2 pl-10 bg-gray-50 border border-gray-200 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
              />
              <Search className="absolute left-3 top-2.5 h-5 w-5 text-gray-400" />
            </div>
          </div>

          <div className="overflow-y-auto h-[calc(100vh-120px)]">
            <DocsNavigation />
          </div>
        </div>
      </div>

      {/* Main Content */}
      <div className="flex-1 min-w-0 overflow-y-auto">
        <div className="max-w-4xl mx-auto px-4 py-8 lg:px-8">
          <DocContent />
        </div>
      </div>

      {/* Right Sidebar - Table of Contents */}
      <div className="hidden xl:block w-64 flex-shrink-0 border-l border-gray-200">
        <div className="sticky top-0 h-screen p-4">
          <h4 className="text-sm font-semibold text-gray-900 mb-4">On this page</h4>
          <div className="overflow-y-auto h-[calc(100vh-80px)]">
            <TableOfContents />
          </div>
        </div>
      </div>
    </div>
  );
};