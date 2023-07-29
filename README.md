# SingleService-Labpro

## :writing_hand: Authors
| Nama                  | NIM      |
| --------------------- | -------- |
| Henry Anand Septian Radityo | 13521004 |

## :running_man: How to run the program
### Prerequisites
- Docker and Docker Compose (for Docker setup)
- Go (for local setup)
### Go Local Setup
1. Clone the repository:
    ```bash
    git clone https://github.com/henryanandsr/SingleService-Labpro
    ```

2. Navigate to the project directory:
    ```bash
    cd SingleService-Labpro/be
    ```

3. Download the dependencies:
    ```bash
    go mod download
    ```

4. Run the application:
    ```bash
    go run main.go
    ```
The application should now be accessible at `http://localhost:3001`.

### Docker setup
1. Clone the repository:
    ```bash
    git clone https://github.com/henryanandsr/SingleService-Labpro
    ```

2. Navigate to the project directory:
    ```bash
    cd SingleService-Labpro/be
    ```

3. Download the dependencies:
    ```bash
   docker-compose up
    ```
The application should now be accessible at `http://localhost:3001`.

### Deployments
1. Server already run at `https://singleservice-labpro-production.up.railway.app`

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
- POST /login
- GET /self
- GET /barang
- GET /barang:id 
- POST /barang
- UPDATE /barang/:id
- DELETE /barang/:id
- GET /perusahaan
- GET /perusahaan/:id
- POST /perusahaan
- UPDATE /perusahaan/:id
- DELETE /perusahaan/:id

## :white_check_mark:	Bonus
1. Deployment
2. Responsive Layout
3. Lighthouse
4. Search Feature

## :bulb: Lighthouse screenshot
<img width="960" alt="katalog" src="https://github.com/henryanandsr/SingleService-Labpro/assets/39207406/e902e5b3-0920-4191-98ec-3c90649ba3da">
<img width="960" alt="home" src = "https://github.com/henryanandsr/SingleService-Labpro/assets/39207406/bfc772f6-de3c-4c42-937f-ca06c6da576f">
<img width="960" alt="login" src = "https://github.com/henryanandsr/SingleService-Labpro/assets/39207406/32938456-1b32-4977-b262-3c106dd771a3">
<img width="960" alt="register" src = "https://github.com/henryanandsr/SingleService-Labpro/assets/39207406/ba5c7fb9-25e0-4ad8-8080-1b82beecccb5">
<img width="960" alt="detailbarang" src="https://github.com/henryanandsr/SingleService-Labpro/assets/39207406/73c45a41-a0f6-4b53-baae-d1644af1544e">
<img width="959" alt="beli" src="https://github.com/henryanandsr/SingleService-Labpro/assets/39207406/3af3ad48-7f03-423c-9bc0-91c7f09f55ea">
<img width="960" alt="riwayat" src="https://github.com/henryanandsr/Monolith-FullStack/assets/39207406/d64254e8-8156-4722-8975-b1daf435daf8">
