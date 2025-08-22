import React, { createContext, useCallback, useContext, useEffect, useMemo, useRef, useState } from 'react'

type ChatMessage = {
  room_id: string
  username: string
  content: string
  created_at: string
}

type WebSocketContextValue = {
  socket: WebSocket | null
  messages: ChatMessage[]
  connect: (roomId: string, clientId: string, username: string) => void
  sendMessage: (content: string) => void
  disconnect: () => void
}

const WebSocketContext = createContext<WebSocketContextValue | undefined>(undefined)
const API_BASE = import.meta.env.VITE_API_URL ?? ''

export const WebSocketProvider: React.FC<React.PropsWithChildren> = ({ children }) => {
  const [socket, setSocket] = useState<WebSocket | null>(null)
  const [messages, setMessages] = useState<ChatMessage[]>([])
  const latestRoomId = useRef<string>('')
  const latestUsername = useRef<string>('')

  const connect = useCallback((roomId: string, clientId: string, username: string) => {
    if (socket) socket.close()
    latestRoomId.current = roomId
    latestUsername.current = username
    const wsUrl = `${(API_BASE || window.location.origin).replace('http', 'ws')}/ws/joinRoom/${roomId}?clientId=${encodeURIComponent(clientId)}&username=${encodeURIComponent(username)}`
    const ws = new WebSocket(wsUrl)
    ws.onopen = () => {
      setMessages([])
    }
    ws.onmessage = (ev) => {
      try {
        const data = JSON.parse(ev.data)
        setMessages((prev) => [...prev, data])
      } catch {
        // ignore
      }
    }
    ws.onclose = () => {
      setSocket(null)
    }
    setSocket(ws)
  }, [socket])

  const sendMessage = useCallback((content: string) => {
    if (!socket || socket.readyState !== WebSocket.OPEN) return
    socket.send(content)
  }, [socket])

  const disconnect = useCallback(() => {
    if (socket) socket.close()
    setSocket(null)
    setMessages([])
  }, [socket])

  const value = useMemo<WebSocketContextValue>(() => ({
    socket,
    messages,
    connect,
    sendMessage,
    disconnect,
  }), [socket, messages, connect, sendMessage, disconnect])

  return (
    <WebSocketContext.Provider value={value}>{children}</WebSocketContext.Provider>
  )
}

export const useWebSocket = () => {
  const ctx = useContext(WebSocketContext)
  if (!ctx) throw new Error('useWebSocket must be used within WebSocketProvider')
  return ctx
}


