'use client'

import { useEffect, useState } from 'react'

// ตัวอย่าง useEffect ที่ต้องมี cleanup
// = หน้านี้สาธิตว่า "เมื่อไหร่ใช้ useEffect / เมื่อไหร่ไม่ใช้"

export default function UseEffectDemo() {
  const [count, setCount] = useState(0)
  const [size, setSize] = useState({ w: 0, h: 0 })

  // ── 1. Interval ที่ต้อง cleanup ──────────────────
  // ใช้ useEffect เพราะต้อง sync กับ external (setInterval)
  useEffect(() => {
    const id = setInterval(() => {
      setCount((c) => c + 1)
    }, 1000)

    // cleanup สำคัญมาก — ไม่งั้น unmount แล้ว interval ยังรัน
    return () => clearInterval(id)
  }, []) // dep ว่าง = รันครั้งเดียวตอน mount

  // ── 2. Window resize event ──────────────────────
  useEffect(() => {
    const handler = () => setSize({ w: window.innerWidth, h: window.innerHeight })
    handler() // initial
    window.addEventListener('resize', handler)
    return () => window.removeEventListener('resize', handler)
  }, [])

  return (
    <main className="max-w-2xl mx-auto p-8 space-y-6">
      <h1 className="text-3xl font-semibold">useEffect demo</h1>

      <section className="bg-white rounded-xl p-5 shadow-sm space-y-2">
        <h2 className="font-semibold">1. Timer (interval)</h2>
        <p className="text-2xl font-mono">{count}s</p>
        <p className="text-sm text-slate-500">
          ใช้ useEffect เพราะต้อง subscribe กับ setInterval · มี cleanup
        </p>
      </section>

      <section className="bg-white rounded-xl p-5 shadow-sm space-y-2">
        <h2 className="font-semibold">2. Window size</h2>
        <p className="font-mono">{size.w} × {size.h}</p>
        <p className="text-sm text-slate-500">
          เปลี่ยนขนาดหน้าจอดูสิ · resize listener + cleanup
        </p>
      </section>

      <section className="bg-amber-50 border border-amber-200 rounded-xl p-5 text-sm text-amber-900">
        <p className="font-semibold mb-2">⚠️ ที่ <strong>ไม่ใช่</strong> useEffect</p>
        <ul className="space-y-1 list-disc list-inside">
          <li>fetch data → ใช้ <code className="bg-white px-1 rounded">useQuery</code> ของ TanStack</li>
          <li>คำนวณค่าจาก state → ใช้ตัวแปรเฉย ๆ (derived value)</li>
          <li>ตอบสนอง user click → ใส่ logic ใน event handler</li>
        </ul>
      </section>
    </main>
  )
}
