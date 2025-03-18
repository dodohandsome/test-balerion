# Back End Take Home # 8

### 1. ติดตั้ง Dependencies
1. รันคำสั่งต่อไปนี้ใน terminal เพื่อดาวน์โหลด dependencies ที่จำเป็น:
```bash
go mod tidy
```

### 2. รันโปรแกรม
1. รันคำสั่ง `go run` เพื่อรันไฟล์ `main.go`
```bash
go run main.go
```


### 3. Run Unit Test
```bash
go test -v 
```

### 4. ได้ Output จากการทำ Unit Test
```plaintext
=== RUN   TestConvertDecimalToThaiBaht
--- PASS: TestConvertDecimalToThaiBaht (0.00s)
PASS
ok      test-backend    0.212s
```