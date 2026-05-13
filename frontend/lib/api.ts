// API client — รวมการเรียก backend ไว้ที่เดียว
// ทุก function throw error ถ้า status ไม่ใช่ 2xx → ให้ TanStack จัด error state

const API_URL = process.env.NEXT_PUBLIC_API_URL ?? 'http://localhost:3000'

export type Major = {
  id: string
  name: string
  code: string
}

export type CreateMajorInput = {
  name: string
  code: string
}

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const res = await fetch(`${API_URL}${path}`, {
    ...init,
    headers: {
      'Content-Type': 'application/json',
      ...(init?.headers ?? {}),
    },
  })
  if (!res.ok) {
    const text = await res.text().catch(() => '')
    throw new Error(`${res.status} ${res.statusText}: ${text}`)
  }
  // 204 No Content
  if (res.status === 204) return undefined as T
  return res.json()
}

// ── Majors ─────────────────────────────────────────────
export const getMajors = () => request<Major[]>('/majors')

export const getMajor = (id: string) => request<Major>(`/majors/${id}`)

export const createMajor = (input: CreateMajorInput) =>
  request<Major>('/majors', {
    method: 'POST',
    body: JSON.stringify(input),
  })

export const deleteMajor = (id: string) =>
  request<void>(`/majors/${id}`, { method: 'DELETE' })
