export const docContent: Record<string, string> = {
  'getting-started': `
    <h1 class="text-3xl font-bold mb-6">Getting Started with Sysmos</h1>
    
    <p class="text-lg mb-6">Welcome to Sysmos! This guide will help you get started with monitoring your servers.</p>
    
    <h2 class="text-2xl font-semibold mb-4">Overview</h2>
    <p class="mb-6">Sysmos provides real-time monitoring, alerting, and analytics for your server infrastructure.</p>
    
    <div class="bg-blue-50 border-l-4 border-blue-500 p-4 mb-6">
      <p class="text-blue-700">Choose a topic from the sidebar to learn more about specific features and capabilities.</p>
    </div>
  `,
  'getting-started/installation': `
    <h1 class="text-3xl font-bold mb-6">Installation Guide</h1>

    <h2 class="text-2xl font-semibold mb-4">System Requirements</h2>
    <ul class="list-disc pl-6 mb-6">
      <li>Linux (Ubuntu 18.04+, CentOS 7+)</li>
      <li>512MB RAM minimum</li>
      <li>1GB disk space</li>
    </ul>

    <h2 class="text-2xl font-semibold mb-4">Installation Steps</h2>
    
    <div class="mb-6">
      <p class="mb-2">1. Download the agent:</p>
      <pre class="bg-gray-800 text-white p-4 rounded"><code>curl -sSL https://serverwatch.io/install.sh | bash</code></pre>
    </div>

    <div class="mb-6">
      <p class="mb-2">2. Configure your API key:</p>
      <pre class="bg-gray-800 text-white p-4 rounded"><code>serverwatch configure --api-key YOUR_API_KEY</code></pre>
    </div>

    <div class="mb-6">
      <p class="mb-2">3. Start the service:</p>
      <pre class="bg-gray-800 text-white p-4 rounded"><code>systemctl start serverwatch</code></pre>
    </div>

    <div class="bg-yellow-50 border-l-4 border-yellow-500 p-4 mt-6">
      <p class="text-yellow-700">Note: Make sure you have root or sudo access before running these commands.</p>
    </div>
  `,
  '404': `
    <div class="text-center py-12">
      <h1 class="text-4xl font-bold text-gray-900 mb-4">Page Not Found</h1>
      <p class="text-xl text-gray-600">The documentation page you're looking for doesn't exist. Please select a topic from the sidebar.</p>
    </div>
  `,
  'core-concepts': `
    <h1 class="text-3xl font-bold mb-6">Core Concepts</h1>

    <h2 class="text-2xl font-semibold mb-4">Architecture Overview</h2>
    <p class="mb-6">ServerWatch uses a distributed architecture with lightweight agents that report to a central server.</p>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
      <div class="bg-white p-6 rounded-lg shadow-md">
        <h3 class="text-xl font-semibold mb-3">Agent</h3>
        <ul class="list-disc pl-6">
          <li>Lightweight process</li>
          <li>Minimal resource usage</li>
          <li>Secure communication</li>
          <li>Auto-updating</li>
        </ul>
      </div>
      <div class="bg-white p-6 rounded-lg shadow-md">
        <h3 class="text-xl font-semibold mb-3">Central Server</h3>
        <ul class="list-disc pl-6">
          <li>Data aggregation</li>
          <li>Analytics processing</li>
          <li>Alert management</li>
          <li>API endpoints</li>
        </ul>
      </div>
    </div>

    <h2 class="text-2xl font-semibold mb-4">Key Features</h2>
    <div class="bg-white p-6 rounded-lg shadow-md mb-6">
      <ul class="space-y-4">
        <li class="flex items-start">
          <span class="inline-flex items-center justify-center h-6 w-6 rounded-full bg-indigo-100 text-indigo-800 mr-3">1</span>
          <div>
            <h4 class="font-semibold">Real-time Monitoring</h4>
            <p class="text-gray-600">Continuous tracking of server metrics with minimal latency</p>
          </div>
        </li>
        <li class="flex items-start">
          <span class="inline-flex items-center justify-center h-6 w-6 rounded-full bg-indigo-100 text-indigo-800 mr-3">2</span>
          <div>
            <h4 class="font-semibold">Intelligent Alerting</h4>
            <p class="text-gray-600">Smart detection of anomalies and potential issues</p>
          </div>
        </li>
        <li class="flex items-start">
          <span class="inline-flex items-center justify-center h-6 w-6 rounded-full bg-indigo-100 text-indigo-800 mr-3">3</span>
          <div>
            <h4 class="font-semibold">Advanced Analytics</h4>
            <p class="text-gray-600">Deep insights into server performance and trends</p>
          </div>
        </li>
      </ul>
    </div>
  `,
};