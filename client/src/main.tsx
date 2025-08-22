import React from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import './index.css'
import { AuthProvider } from './state/AuthContext'
import { WebSocketProvider } from './/state/WebSocketContext'
import LoginPage from './pages/LoginPage'
import HomePage from './/pages/HomePage'
import AppPage from './/pages/AppPage'
import RootLayout from './layouts/RootLayout'
import Guard from './components/Guard'

const router = createBrowserRouter([
  { path: '/', element: <RootLayout><Guard><HomePage /></Guard></RootLayout> },
  { path: '/login', element: <RootLayout><LoginPage /></RootLayout> },
  { path: '/app/:roomId', element: <RootLayout><Guard><AppPage /></Guard></RootLayout> },
])

createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <AuthProvider>
      <WebSocketProvider>
        <RouterProvider router={router} />
      </WebSocketProvider>
    </AuthProvider>
  </React.StrictMode>
)


