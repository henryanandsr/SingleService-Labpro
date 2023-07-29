# SingleService-Labpro

## :writing_hand: Authors
| Nama                  | NIM      |
| --------------------- | -------- |
| Henry Anand Septian Radityo | 13521004 |

## :running_man: How to run the program
### Local
1. `go run main.go` or `docker-compose up`
### Deployments
1. Already ran at `https://singleservice-labpro-production.up.railway.app`

## :books: Design Pattern
1. Singleton Pattern
   Singleton dipakai untuk menginstansiasi sebuah database. Singleton pattern akan meminimalisir penggunaan resource secara berlebihan karena instance yang dibuat pada database hanya akan dilakukan sekali.
3. Repository Pattern
   Repository pattern adalah pattern untuk memisahkan antara algoritma logic dan manipulasi database. Penggunaan pattern ini akan membuat de-coupling antara data dan bisnis logic sehingga memudahkan untuk memaintain data.
5. MVC Pattern
   Pola model-view-controller akan membuat kode lebih teroganisir dengan memisah - misahkan berdasarkan tugas masing - masing. Model untuk data dan logika bisnis, view untuk tampilan, dan controller untuk mengatur interaksi keduanya.

## :wrench: Tech Stack
Go (Golang), Postgresql, Gin, GORM, Docker

## :purple_circle: Endpoint
POST /login
GET /self
GET /barang
GET /barang:id 
POST /barang
UPDATE /barang/:id
DELETE /barang/:id
GET /perusahaan
GET /perusahaan/:id
POST /perusahaan
UPDATE /perusahaan/:id
DELETE /perusahaan/:id

## :white_check_mark:	Bonus
1. Deployment
