ใช้ Echo สร้าง REST API
แยกโฟลเดอร์แบบ Hexagonal (domain, app, infra, config, cmd)
Logging ด้วย slog
โหลด config จาก env.yaml ด้วย viper
เชื่อมต่อ SQL Server สำหรับเก็บ Order
สร้าง memory cache สำหรับ product (refresh ทุก 5 นาที)
เชื่อมต่อ external API (http://pushnoti.com/api/test)
มี unit test สำหรับ handler หลัก
วิธีใช้งานเบื้องต้น:
config อยู่ที่ internal/config/env.yaml
main entry point: cmd/api/main.go
สั่งรัน: go run ./cmd/api
สั่งเทส: go test ./internal/app/order