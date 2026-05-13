import Link from 'next/link'

export default function Home() {
  return (
    <main className="max-w-2xl mx-auto p-8 space-y-6">
      <h1 className="text-3xl font-semibold">HLLC Workshop · Day 3</h1>
      <p className="text-slate-600">หน้า demo การเชื่อม FE กับ BE Major + Course ที่เขียนเมื่อวาน</p>

      <ul className="space-y-2">
        <li>
          <Link href="/majors" className="text-blue-600 hover:underline">
            → หน้า /majors (useQuery + useMutation)
          </Link>
        </li>
        <li>
          <Link href="/useeffect-demo" className="text-blue-600 hover:underline">
            → หน้า /useeffect-demo (useEffect + cleanup)
          </Link>
        </li>
      </ul>
    </main>
  )
}
