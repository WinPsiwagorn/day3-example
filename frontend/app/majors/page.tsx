'use client'

import { useState } from 'react'
import { useCreateMajor, useDeleteMajor, useMajors } from '@/lib/hooks/use-majors'
import { useUIStore } from '@/stores/ui'

export default function MajorsPage() {
  const { data, isLoading, error } = useMajors()
  const createMutation = useCreateMajor()
  const deleteMutation = useDeleteMajor()
  const pushToast = useUIStore((s) => s.pushToast)

  const [name, setName] = useState('')
  const [code, setCode] = useState('')

  async function handleCreate(e: React.FormEvent) {
    e.preventDefault()
    if (!name || !code) return
    createMutation.mutate(
      { name, code },
      {
        onSuccess: () => {
          setName('')
          setCode('')
          pushToast(`สร้าง ${name} สำเร็จ`)
        },
        onError: (err) => pushToast(`พัง: ${err.message}`),
      },
    )
  }

  return (
    <main className="max-w-3xl mx-auto p-8 space-y-8">
      <header>
        <h1 className="text-3xl font-semibold">Majors</h1>
        <p className="text-slate-600">list + create · ใช้ TanStack Query</p>
      </header>

      {/* ── form ── */}
      <form onSubmit={handleCreate} className="bg-white rounded-xl p-5 shadow-sm space-y-3">
        <h2 className="font-semibold">เพิ่ม Major ใหม่</h2>
        <div className="grid grid-cols-2 gap-3">
          <input
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="ชื่อ เช่น Software Engineering"
            className="border rounded px-3 py-2 text-sm"
          />
          <input
            value={code}
            onChange={(e) => setCode(e.target.value)}
            placeholder="code เช่น SE"
            className="border rounded px-3 py-2 text-sm"
          />
        </div>
        <button
          type="submit"
          disabled={createMutation.isPending}
          className="bg-slate-900 text-white rounded px-4 py-2 text-sm disabled:opacity-50"
        >
          {createMutation.isPending ? 'กำลังบันทึก...' : 'บันทึก'}
        </button>
      </form>

      {/* ── list ── */}
      <section className="bg-white rounded-xl p-5 shadow-sm">
        <h2 className="font-semibold mb-3">รายการ</h2>

        {isLoading && <p className="text-slate-500">กำลังโหลด...</p>}
        {error && <p className="text-rose-600">โหลดไม่สำเร็จ: {error.message}</p>}

        {data && data.length === 0 && (
          <p className="text-slate-500">ยังไม่มีข้อมูล · เพิ่มอันแรกเลย</p>
        )}

        {data && data.length > 0 && (
          <ul className="divide-y">
            {data.map((m) => (
              <li key={m.id} className="py-3 flex items-center justify-between">
                <div>
                  <p className="font-medium">{m.name}</p>
                  <p className="text-xs text-slate-500 font-mono">{m.code}</p>
                </div>
                <button
                  onClick={() => deleteMutation.mutate(m.id)}
                  className="text-sm text-rose-600 hover:underline"
                >
                  ลบ
                </button>
              </li>
            ))}
          </ul>
        )}
      </section>

      <ToastList />
    </main>
  )
}

function ToastList() {
  const toasts = useUIStore((s) => s.toasts)
  const removeToast = useUIStore((s) => s.removeToast)

  return (
    <div className="fixed bottom-4 right-4 space-y-2">
      {toasts.map((t) => (
        <div
          key={t.id}
          className="bg-slate-900 text-white text-sm px-4 py-2 rounded shadow cursor-pointer"
          onClick={() => removeToast(t.id)}
        >
          {t.text}
        </div>
      ))}
    </div>
  )
}
