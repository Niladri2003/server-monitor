import React, { useState } from 'react';
import { Activity, Menu, X } from 'lucide-react';

export default function Navbar() {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  return (
    <nav className="fixed w-full bg-white/90 backdrop-blur-sm z-50 shadow-sm">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          <div className="flex items-center space-x-2">
            <Activity className="h-8 w-8 text-indigo-600" />
            <span className="text-xl font-bold text-gray-900">Sysmos</span>
          </div>
          
          {/* Desktop Menu */}
          <div className="hidden md:block">
            <div className="flex items-center space-x-8">
              <a href="#features" className="text-gray-700 hover:text-indigo-600 transition-colors">Features</a>
              <a href="#metrics" className="text-gray-700 hover:text-indigo-600 transition-colors">Analytics</a>
              <a href="#insights" className="text-gray-700 hover:text-indigo-600 transition-colors">Insights</a>
              <button className="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition-colors">
                Get Started
              </button>
            </div>
          </div>

          {/* Mobile Menu Button */}
          <div className="md:hidden">
            <button
              onClick={() => setIsMenuOpen(!isMenuOpen)}
              className="text-gray-700 hover:text-indigo-600 transition-colors"
            >
              {isMenuOpen ? <X className="h-6 w-6" /> : <Menu className="h-6 w-6" />}
            </button>
          </div>
        </div>

        {/* Mobile Menu */}
        {isMenuOpen && (
          <div className="md:hidden py-4 border-t border-gray-100">
            <div className="flex flex-col space-y-4">
              <a 
                href="#features" 
                className="text-gray-700 hover:text-indigo-600 transition-colors px-4"
                onClick={() => setIsMenuOpen(false)}
              >
                Features
              </a>
              <a 
                href="#metrics" 
                className="text-gray-700 hover:text-indigo-600 transition-colors px-4"
                onClick={() => setIsMenuOpen(false)}
              >
                Analytics
              </a>
              <a 
                href="#insights" 
                className="text-gray-700 hover:text-indigo-600 transition-colors px-4"
                onClick={() => setIsMenuOpen(false)}
              >
                Insights
              </a>
              <div className="px-4 pb-2">
                <button className="w-full bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition-colors">
                  Get Started
                </button>
              </div>
            </div>
          </div>
        )}
      </div>
    </nav>
  );
}