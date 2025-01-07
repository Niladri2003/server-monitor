
import Navbar from './components/Navbar';
import Hero from './components/Hero';
import Features from './components/Features';
import GlobalMonitoring from './components/GlobalMonitoring';
import MetricsDashboard from './components/MetricsDashboard';
import Footer from './components/Footer';

function App() {
  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />
      <Hero />
      <Features />
      <GlobalMonitoring />
      <MetricsDashboard />
      <Footer />
    </div>
  );
}

export default App;