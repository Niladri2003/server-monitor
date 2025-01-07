import React from 'react';
import { FeatureCard } from './FeatureCard';
import { features } from '../data/features';

export default function Features() {
  return (
    <section id="features" className="py-16 md:py-20">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center mb-12 md:mb-16">
          <h2 className="text-2xl md:text-3xl font-bold text-gray-900">Comprehensive Monitoring Features</h2>
          <p className="mt-4 text-lg md:text-xl text-gray-600">Everything you need to keep your servers running smoothly</p>
        </div>
        
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 md:gap-8">
          {features.map((feature, index) => (
            <FeatureCard key={index} {...feature} />
          ))}
        </div>
      </div>
    </section>
  );
}