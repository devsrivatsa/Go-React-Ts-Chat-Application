import React, { useEffect, useState } from 'react'

const themes = ['pastel', 'emerald', 'synthwave', 'cupcake', 'garden']

const ThemeToggle: React.FC = () => {
  const [theme, setTheme] = useState<string>(() => localStorage.getItem('theme') || 'pastel')
  useEffect(() => {
    document.documentElement.setAttribute('data-theme', theme)
    localStorage.setItem('theme', theme)
  }, [theme])
  return (
    <select className="select select-bordered" value={theme} onChange={(e)=>setTheme(e.target.value)}>
      {themes.map(t => <option key={t} value={t}>{t}</option>)}
    </select>
  )
}

export default ThemeToggle


