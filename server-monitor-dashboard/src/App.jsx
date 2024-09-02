import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import {SidebarDemo} from "./Pages/Sidebar.jsx";
import {Route, Router, Routes} from "react-router-dom";
import Dashboard from "./Pages/Dashboard.jsx";
import {Signup} from "./Pages/Signup.jsx";
import {Login} from "./Pages/Login.jsx";



function App() {

  return (

          <Routes>
              <Route path="/" element={<Signup/>}/>
              <Route path="login" element={<Login/>}/>
              <Route path="/dashboard" element={<SidebarDemo />}>
                  <Route path="admin" element={<Dashboard />} />
              </Route>
          </Routes>

  )
}

export default App
