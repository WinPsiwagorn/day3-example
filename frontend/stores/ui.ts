import { create } from 'zustand'

// ตัวอย่าง Zustand store — client state ที่แชร์ทั่ว app
// ใช้กับ state ที่ "ไม่เกี่ยวกับ server data" เช่น modal เปิด/ปิด, theme, toast queue
//
// data ที่มาจาก API → ใช้ TanStack Query แทน (ดู use-majors.ts)

type Toast = { id: number; text: string }

type UIStore = {
  // modal "create major"
  createOpen: boolean
  openCreate: () => void
  closeCreate: () => void

  // toast list
  toasts: Toast[]
  pushToast: (text: string) => void
  removeToast: (id: number) => void
}

export const useUIStore = create<UIStore>((set) => ({
  createOpen: false,
  openCreate: () => set({ createOpen: true }),
  closeCreate: () => set({ createOpen: false }),

  toasts: [],
  pushToast: (text) =>
    set((s) => ({ toasts: [...s.toasts, { id: Date.now(), text }] })),
  removeToast: (id) =>
    set((s) => ({ toasts: s.toasts.filter((t) => t.id !== id) })),
}))
