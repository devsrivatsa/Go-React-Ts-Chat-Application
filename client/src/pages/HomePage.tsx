import React, { useEffect, useMemo, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { useAuth } from '../state/AuthContext'
import { nanoid } from 'nanoid'
import ThemeToggle from '../components/ThemeToggle'
import RoomListItem from '../components/RoomListItem'

const API_BASE = import.meta.env.VITE_API_URL ?? ''

type Room = { id: string; name: string }

const HomePage: React.FC = () => {
  const { user, isAuthenticated, logout } = useAuth()
  const navigate = useNavigate()
  const [rooms, setRooms] = useState<Room[]>([])
  const [newRoomName, setNewRoomName] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    (async () => {
      const res = await fetch(`${API_BASE}/ws/getRooms`)
      if (res.ok) {
        setRooms(await res.json())
      }
    })()
  }, [])

  const createRoom = async () => {
    if (!newRoomName) return
    setLoading(true)
    setError(null)
    try {
      const id = nanoid(8)
      const res = await fetch(`${API_BASE}/ws/createRoom`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ id, name: newRoomName })
      })
      if (!res.ok) throw new Error('Failed to create room')
      const created = await res.json()
      setRooms((prev) => [...prev, created])
      setNewRoomName('')
    } catch (e: any) {
      setError(e.message)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="min-h-screen bg-base-200">
      <div className="navbar bg-base-100 shadow">
        <div className="flex-1">
          <a className="btn btn-ghost text-xl">Chat Rooms</a>
        </div>
        <div className="flex-none gap-2">
          <ThemeToggle />
          {isAuthenticated ? (
            <>
              <div className="badge badge-primary">{user?.username}</div>
              <button className="btn btn-ghost" onClick={logout}>Logout</button>
            </>
          ) : (
            <Link className="btn btn-primary" to="/login">Login</Link>
          )}
        </div>
      </div>

      <div className="container mx-auto p-4 grid gap-6 md:grid-cols-3">
        <div className="md:col-span-2">
          <h2 className="text-xl font-semibold mb-3">Available Rooms</h2>
          <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-4">
            {rooms.map((r) => (
              <RoomListItem key={r.id} id={r.id} name={r.name} onJoin={() => navigate(`/app/${r.id}`)} />
            ))}
            {rooms.length === 0 && <div className="opacity-50">No rooms yet. Create one!</div>}
          </div>
        </div>
        <div>
          <div className="card bg-base-100 shadow">
            <div className="card-body">
              <h2 className="card-title">Create a Room</h2>
              {error && <div className="alert alert-error text-sm">{error}</div>}
              <input className="input input-bordered" placeholder="Room name" value={newRoomName} onChange={e=>setNewRoomName(e.target.value)} />
              <button className={`btn btn-primary ${loading ? 'loading' : ''}`} onClick={createRoom}>Create</button>
            </div>
          </div>
        </div>
      </div>  
    </div>
  )
}

export default HomePage


