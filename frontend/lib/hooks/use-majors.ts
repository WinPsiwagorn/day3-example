'use client'

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { createMajor, deleteMajor, getMajor, getMajors } from '@/lib/api'

// ── query hooks ─────────────────────────────────────────
// queryKey เป็น array — ใช้แบบ ['majors'] หรือ ['majors', id]
// ถ้า invalidateQueries(['majors']) จะ refetch ทุก hook ที่ขึ้นต้นด้วย ['majors']

export function useMajors() {
  return useQuery({
    queryKey: ['majors'],
    queryFn: getMajors,
  })
}

export function useMajor(id: string) {
  return useQuery({
    queryKey: ['majors', id],
    queryFn: () => getMajor(id),
    enabled: !!id,
  })
}

// ── mutation hooks ──────────────────────────────────────
// invalidate cache หลัง success → list refetch ใหม่อัตโนมัติ

export function useCreateMajor() {
  const qc = useQueryClient()
  return useMutation({
    mutationFn: createMajor,
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['majors'] })
    },
  })
}

export function useDeleteMajor() {
  const qc = useQueryClient()
  return useMutation({
    mutationFn: deleteMajor,
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['majors'] })
    },
  })
}
