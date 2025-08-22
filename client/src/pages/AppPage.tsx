import React, { useEffect, useMemo, useState } from 'react'
import { Link, useNavigate, useParams } from 'react-router-dom'
import { useWebSocket } from '../state/WebSocketContext'
import { useAuth } from '../state/AuthContext'
import { nanoid } from 'nanoid'

const AppPage: React.FC = () => {
  const { roomId = '' } = useParams()
  const { user, isAuthenticated } = useAuth()
  const { messages, connect, disconnect, sendMessage } = useWebSocket()
  const navigate = useNavigate()

  const [draft, setDraft] = useState('')

  useEffect(() => {
    const clientId = user?.id ?? nanoid(6)
    const username = user?.username ?? `guest-${clientId}`
    connect(roomId, clientId, username)
    return () => disconnect()
  }, [roomId])

  const onSend = () => {
    if (!draft.trim()) return
    sendMessage(draft.trim())
    setDraft('')
  }

  return (
    <div className="min-h-screen bg-base-200 flex flex-col">
      <div className="navbar bg-base-100 shadow">
        <div className="flex-1">
          <Link className="btn btn-ghost text-xl" to="/">← Rooms</Link>
        </div>
        <div className="flex-none gap-2">
          <div className="badge badge-primary">Room {roomId}</div>
          {isAuthenticated && <div className="badge badge-ghost">{user?.username}</div>}
        </div>
      </div>
      <div className="container mx-auto flex-1 p-4 grid grid-rows-[1fr_auto] max-w-4xl w-full">
        <div className="bg-base-100 rounded-xl shadow p-4 overflow-y-auto space-y-3">
          {messages.map((m, idx) => (
            <div key={idx} className="chat chat-start">
              <div className="chat-header">{m.username}<time className="text-xs opacity-50 ml-2">{new Date(m.created_at).toLocaleTimeString()}</time></div>
              <div className="chat-bubble chat-bubble-secondary whitespace-pre-wrap">{m.content}</div>
            </div>
          ))}
          {messages.length === 0 && <div className="opacity-60">No messages yet. Say hi!</div>}
        </div>
        <div className="mt-3 flex gap-2">
          <input className="input input-bordered flex-1" placeholder="Type a message..." value={draft} onChange={e=>setDraft(e.target.value)} onKeyDown={e=>{if(e.key==='Enter') onSend()}} />
          <button className="btn btn-primary" onClick={onSend}>Send</button>
        </div>
      </div>
    </div>
  )
}

export default AppPage


