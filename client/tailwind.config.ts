import type { Config } from 'tailwindcss'

export default {
  content: [
    './index.html',
    './src/**/*.{ts,tsx}',
  ],
  theme: {
    extend: {},
  },
  plugins: [require('daisyui')],
  daisyui: {
    themes: ['pastel', 'emerald', 'synthwave', 'cupcake', 'garden'],
    darkTheme: 'synthwave'
  }
} satisfies Config


