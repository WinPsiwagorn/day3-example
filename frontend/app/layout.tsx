import type { Metadata } from 'next'
import { Providers } from './providers'
import './globals.css'

export const metadata: Metadata = {
  title: 'HLLC Workshop · Day 3',
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="th">
      <body className="min-h-screen text-slate-900">
        <Providers>{children}</Providers>
      </body>
    </html>
  )
}
