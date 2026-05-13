# HLLC Workshop · Day 3 — Example code

โปรเจกต์ตัวอย่างสำหรับ Day 3 มี 2 ฝั่ง:

- **`backend/`** — Go + Fiber v3 + MongoDB · เป็น **reference** ของ BE ที่ทำใน Day 2 (Major + Course CRUD แบบ layered)
- **`frontend/`** — Next.js 15 + Tailwind + TanStack Query + Zustand · ตัวอย่าง Day 3 topics

---

## 1. รัน Backend

```bash
cd backend
cp .env.example .env       # แก้ MONGO_URI ถ้าจำเป็น
go mod tidy
go run main.go
```

ขึ้นที่ <http://localhost:3000> · ลอง:

```bash
curl http://localhost:3000/health
# {"ok":true}

curl -X POST http://localhost:3000/majors \
  -H "Content-Type: application/json" \
  -d '{"name":"Software Engineering","code":"SE"}'

curl http://localhost:3000/majors
```

### โครง backend

```
backend/
├── main.go              # wire ทุก layer + CORS + Listen
├── database/mongo.go    # connect Mongo + Ping
├── major/               # Major CRUD (โจทย์เช้า Day 2)
│   ├── model.go         # type Major + bson tags
│   ├── dto.go           # CreateMajorDTO + UpdateMajorDTO + validate tags
│   ├── repository.go    # คุย Mongo: Insert / FindAll / FindByID / Update / Delete
│   ├── service.go       # business logic · ไม่รู้จัก Fiber, ไม่รู้จัก Mongo
│   └── handler.go       # HTTP layer + RegisterRoutes
└── course/              # Course CRUD (โจทย์บ่าย Day 2)
    └── ... โครงเดียวกัน + 3 challenge:
            int validation, []ObjectID, nested struct
```

### Endpoints

| Method | Path | คำอธิบาย |
|---|---|---|
| GET | `/health` | health check |
| GET | `/majors` | list ทั้งหมด |
| GET | `/majors/:id` | get by id |
| POST | `/majors` | create |
| PUT | `/majors/:id` | update (partial) |
| DELETE | `/majors/:id` | delete |
| GET | `/courses` | list (รองรับ `?major=<id>`) |
| GET | `/courses/:id` | get by id |
| POST | `/courses` | create |
| PUT | `/courses/:id` | update |
| DELETE | `/courses/:id` | delete |

---

## 2. รัน Frontend

ต้อง **เปิด backend ก่อน** เพราะ FE เรียก `localhost:3000`

```bash
cd frontend
cp .env.local.example .env.local
pnpm install
pnpm dev
```

ขึ้นที่ <http://localhost:3001>

### หน้าที่มีตัวอย่าง

| Path | สอนเรื่อง |
|---|---|
| `/` | หน้าแรก · ลิงค์ไปแต่ละ demo |
| `/majors` | **TanStack Query** (useQuery + useMutation) + **Zustand** (toast store) |
| `/useeffect-demo` | **useEffect + cleanup** (interval, window resize) |

### โครง frontend

```
frontend/
├── app/
│   ├── layout.tsx           # ครอบ <Providers>
│   ├── providers.tsx        # QueryClient + Devtools
│   ├── page.tsx             # home
│   ├── majors/page.tsx      # ตัวอย่าง TanStack + Zustand
│   └── useeffect-demo/...   # ตัวอย่าง useEffect
├── lib/
│   ├── api.ts               # API client · fetch wrapper + types
│   └── hooks/
│       └── use-majors.ts    # custom hook ที่ห่อ useQuery/useMutation
└── stores/
    └── ui.ts                # Zustand store · modal + toast
```

---

## 3. CORS

Backend เปิดให้ origin `http://localhost:3001` เรียกได้ (ดู `backend/main.go`)

ถ้า FE รันที่ port อื่น → แก้ `ALLOWED_ORIGIN` ใน `backend/.env`

---

## 4. ลำดับที่แนะนำให้น้อง

1. **clone repo** · เปิด 2 terminal
2. **terminal 1**: รัน backend · ลอง `/majors` ผ่าน Postman
3. **terminal 2**: รัน frontend · เปิด `/majors`
4. ลองกด "เพิ่ม Major" บน FE → ดู Network tab → ดู cache update
5. เปิด **TanStack Devtools** (มุมล่างซ้าย) ดู query state real-time
6. อ่าน code ทีละไฟล์ตาม comment · เริ่มจาก `lib/api.ts` → `lib/hooks/use-majors.ts` → `app/majors/page.tsx`
