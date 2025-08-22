import React from 'react'
import { Navigate } from 'react-router-dom'
import { useAuth } from '../state/AuthContext'

const Guard: React.FC<React.PropsWithChildren> = ({ children }) => {
  const { isAuthenticated } = useAuth()
  if (!isAuthenticated) return <Navigate to="/login" replace />
  return <>{children}</>
}

export default Guard


