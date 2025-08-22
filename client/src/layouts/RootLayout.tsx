import React from 'react'

const RootLayout: React.FC<React.PropsWithChildren> = ({ children }) => {
  return (
    <div className="min-h-screen" data-theme={localStorage.getItem('theme') || 'pastel'}>
      {children}
    </div>
  )
}

export default RootLayout


