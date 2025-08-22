import React, { createContext, useCallback, useContext, useEffect, useMemo, useState } from 'react'

type User = {
  id: string
  username: string
  email?: string
}

type AuthContextValue = {
  user: User | null
  isAuthenticated: boolean
  login: (email: string, password: string) => Promise<void>
  signup: (username: string, email: string, password: string) => Promise<void>
  logout: () => Promise<void>
}

const AuthContext = createContext<AuthContextValue | undefined>(undefined)

const API_BASE = import.meta.env.VITE_API_URL ?? ''

export const AuthProvider: React.FC<React.PropsWithChildren> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null)

  const login = useCallback(async (email: string, password: string) => {
    const res = await fetch(`${API_BASE}/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      body: JSON.stringify({ email, password }),
    })
    if (!res.ok) {
      throw new Error('Login failed')
    }
    const data = await res.json()
    setUser({ id: data.id, username: data.username, email: data.email })
  }, [])

  const signup = useCallback(async (username: string, email: string, password: string) => {
    const res = await fetch(`${API_BASE}/signup`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, email, password }),
    })
    if (!res.ok) {
      throw new Error('Signup failed')
    }
    // Auto-login after signup
    await login(email, password)
  }, [login])

  const logout = useCallback(async () => {
    await fetch(`${API_BASE}/logout`, { credentials: 'include' })
    setUser(null)
  }, [])

  const value = useMemo<AuthContextValue>(() => ({
    user,
    isAuthenticated: !!user,
    login,
    signup,
    logout,
  }), [user, login, signup, logout])

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}

export const useAuth = () => {
  const ctx = useContext(AuthContext)
  if (!ctx) throw new Error('useAuth must be used within AuthProvider')
  return ctx
}


