
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { Layout } from './components/layout/Layout';
import { HomePage } from './pages/HomePage';
import { LoginPage } from './pages/auth/LoginPage';
import { SignupPage } from './pages/auth/SignupPage';
import { DashboardPage } from './pages/DashboardPage';
import { ServersPage } from './pages/ServersPage';
import { ApiKeysPage } from './pages/ApiKeysPage';
import { AlertsPage } from './pages/AlertsPage';
import { TeamsPage } from './pages/TeamsPage';
import { SettingsPage } from './pages/SettingsPage';
import { DocumentationPage } from './pages/DocumentationPage';
import { PortMonitoringPage } from './pages/PortMonitoringPage';

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
{/*         <Route path="/" element={<HomePage />} /> */}
        <Route path="/login" element={<LoginPage />} />
        <Route path="/signup" element={<SignupPage />} />
        <Route path="/docs/*" element={<DocumentationPage />} />
        <Route path="/" element={<Layout />}>
          <Route path="dashboard" element={<DashboardPage />} />
          <Route path="servers" element={<ServersPage />} />
          <Route path="port-monitoring" element={<PortMonitoringPage />} />
          <Route path="api-keys" element={<ApiKeysPage />} />
          <Route path="alerts" element={<AlertsPage />} />
          <Route path="teams" element={<TeamsPage />} />
          <Route path="settings" element={<SettingsPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}
