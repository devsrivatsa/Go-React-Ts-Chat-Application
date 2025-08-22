import React, { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { useAuth } from '../state/AuthContext'

const LoginPage: React.FC = () => {
  const { login, signup } = useAuth()
  const navigate = useNavigate()
  const [mode, setMode] = useState<'login' | 'signup'>('login')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [username, setUsername] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const onSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setLoading(true)
    setError(null)
    try {
      if (mode === 'login') {
        await login(email, password)
      } else {
        await signup(username, email, password)
      }
      navigate('/')
    } catch (err: any) {
      setError(err.message ?? 'Something went wrong')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="min-h-screen bg-base-200 flex items-center justify-center p-4">
      <div className="card w-full max-w-md shadow-xl bg-base-100">
        <div className="card-body">
          <h2 className="card-title text-2xl mb-2">{mode === 'login' ? 'Welcome back' : 'Create your account'}</h2>
          <p className="opacity-70 mb-4">Chat Rooms • Real-time • Pretty UI</p>
          {error && <div className="alert alert-error text-sm mb-2">{error}</div>}
          <form onSubmit={onSubmit} className="space-y-3">
            {mode === 'signup' && (
              <div className="form-control">
                <label className="label"><span className="label-text">Username</span></label>
                <input className="input input-bordered" value={username} onChange={e=>setUsername(e.target.value)} required/>
              </div>
            )}
            <div className="form-control">
              <label className="label"><span className="label-text">Email</span></label>
              <input type="email" className="input input-bordered" value={email} onChange={e=>setEmail(e.target.value)} required/>
            </div>
            <div className="form-control">
              <label className="label"><span className="label-text">Password</span></label>
              <input type="password" className="input input-bordered" value={password} onChange={e=>setPassword(e.target.value)} required/>
            </div>
            <div className="form-control mt-4">
              <button className={`btn btn-primary ${loading ? 'loading' : ''}`} type="submit">
                {mode === 'login' ? 'Login' : 'Sign up'}
              </button>
            </div>
          </form>
          <div className="divider">or</div>
          <button className="btn btn-ghost" onClick={() => setMode(mode === 'login' ? 'signup' : 'login')}>
            {mode === 'login' ? 'Create an account' : 'Have an account? Log in'}
          </button>
        </div>
      </div>
    </div>
  )
}

export default LoginPage


