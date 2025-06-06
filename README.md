# 006-go-api-orm-gorm

โปรเจกต์ตัวอย่าง Go API ที่ใช้ [Gin](https://github.com/gin-gonic/gin) เป็นเว็บเฟรมเวิร์ก และ [GORM](https://gorm.io/) เป็น ORM สำหรับจัดการฐานข้อมูล PostgreSQL

## 💻 เริ่มต้นสร้างโปรเจกต์

1. สร้างโฟลเดอร์โปรเจกต์

```bash
mkdir 006-go-api-orm-gorm
cd 006-go-api-orm-gorm
```

2. สร้างโมดูล Go

```bash
go mod init 006-go-api-orm-gorm
```

3. ติดตั้งไลบรารีที่จำเป็น

```bash
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/joho/godotenv
```

## 📂 โครงสร้างไฟล์โปรเจกต์

```bash
006-go-api-orm-gorm/
├── config/
│   └── config.go       # โหลด config DB จาก .env หรือ ENV
├── controllers/       # จัดการ logic ควบคุม API
├── database/
│   └── database.go    # เชื่อมต่อฐานข้อมูลและจัดการ migration
├── models/            # โครงสร้างข้อมูลฐานข้อมูล (Model)
├── routes/            # กำหนดเส้นทาง API
├── main.go            # จุดเริ่มต้นของแอป
├── go.mod
├── go.sum
├── README.md
└── .env               # กำหนดค่าคอนฟิกฐานข้อมูล
```

## ⚙️ การตั้งค่า Database

สร้างไฟล์ .env ในโฟลเดอร์โปรเจกต์:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=goorm
DB_PORT=5432
DB_SSLMODE=disable
```

## 🚀 รันโปรเจกต์

1. เชื่อมต่อฐานข้อมูล PostgreSQL ให้เรียบร้อย (สร้างฐานข้อมูล goorm)
2. รันคำสั่ง

```bash
go run main.go
```

3. เปิดเบราว์เซอร์ไปที่

```bash
http://localhost:8080/
```

จะเห็นข้อความ

```json
{ "message": "Welcome to 006-go-api-orm-gorm" }
```

## 📋 คำอธิบาย

- ใช้ Gin สำหรับ HTTP routing และ middleware
- ใช้ GORM สำหรับ ORM จัดการฐานข้อมูล PostgreSQL
- โหลด config DB จากไฟล์ .env หรือ environment variables ด้วย github.com/joho/godotenv
- โครงสร้างแยกชัดเจนตามหน้าที่

## 📚 เรียนรู้เพิ่มเติม

- [Gin Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [godotenv Documentation](https://github.com/joho/godotenv)

📧 ติดต่อ
ถ้ามีคำถามหรือข้อเสนอแนะ ติดต่อได้เลยครับ!
---

📄 License  
MIT License – ใช้งานฟรีเพื่อการศึกษาและพัฒนาต่อได้เต็มที่

🙋‍♂️ ผู้พัฒนา  
พัฒนาโดยคุณ [ปลื้ม ชินวัตร]  
📫 ติดต่อ: chinawat.dc@gmail.com
