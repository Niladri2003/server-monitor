
import { MetricsPreview } from './MetricsPreview';

export default function Hero() {
  return (
    <section className="pt-16 md:pt-20 pb-24 md:pb-32 bg-gradient-to-br from-indigo-900 via-indigo-800 to-indigo-900 text-white">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center space-y-6 md:space-y-8 py-12 md:py-20">
          <h1 className="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight animate-fade-in px-4">
            Intelligent System
            <span className="text-indigo-400"> Monitoring</span>
          </h1>
          <p className="text-lg md:text-xl text-gray-300 max-w-2xl mx-auto px-4">
            Sysmos provides the complete set of tools needed to connect and process data of your Production server in order to  real-time metrics, and predictive analytics for your infrastructure. Utilise Effectively your Resource using Sysmos
          </p>
          <div className="flex flex-col sm:flex-row justify-center space-y-4 sm:space-y-0 sm:space-x-4 px-4">
            <button className="bg-indigo-600 text-white px-6 md:px-8 py-3 rounded-lg hover:bg-indigo-700 transition-colors w-full sm:w-auto">
              Start Monitoring
            </button>
            <button className="bg-white text-indigo-900 px-6 md:px-8 py-3 rounded-lg hover:bg-gray-100 transition-colors w-full sm:w-auto">
              Watch Demo
            </button>
          </div>
        </div>
        <div className="px-2 sm:px-4">
          <MetricsPreview />
        </div>
      </div>
    </section>
  );
}