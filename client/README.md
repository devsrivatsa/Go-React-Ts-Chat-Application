Chat App Client (Vite + React 19 + Tailwind + DaisyUI)

Getting Started

- Install Node 24.6+ and pnpm/npm/yarn
- Install dependencies and run dev server:

```bash
pnpm i # or npm i / yarn
pnpm dev # or npm run dev / yarn dev
```

Server is expected on http://localhost:8080. WebSocket path: `ws://localhost:8080/ws/joinRoom/:roomId?clientId=...&username=...`.

You can change the backend base via `VITE_API_BASE` env var.






