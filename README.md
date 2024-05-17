# TreckTicket

## About TreckTicket
Aplikasi pembelian tiket wisata dirancang untuk memenuhi kebutuhan praktis dalam merencanakan perjalanan dan membeli tiket secara mudah. Dalam era digital ini, aplikasi tersebut menyediakan akses yang nyaman bagi pengguna untuk memesan tiket wisata melalui perangkat seluler mereka. Dengan kemudahan ini, para pelancong dapat menghindari proses konvensional yang memakan waktu dan tenaga. Aplikasi ini juga dapat memberikan rekomendasi destinasi wisata dan menawarkan berbagai pilihan tempat wisata. Dengan demikian, aplikasi ini tidak hanya memberikan kemudahan dalam pembelian tiket, tetapi juga meningkatkan pengalaman wisata pengguna dengan menyediakan informasi yang relevan dan pilihan yang sesuai dengan kebutuhan mereka.

## Features
### User
- User BisaLogin
- User Bisa Register
- User Bisa Mencari Tempat Wisata
- User Bisa Melakukan Pembelian Tiket
- User Bisa Melihat Tiket Yang Sudah Dibeli

### Admin
- Admin Bisa Login
- Admin Bisa Mengedit Informasi Tempat Wisata Yang Dia Kelola
- Admin Bisa Melihat List Transaksi Pembelian Tiket di Tempat Wisata Yang Dia Kelola


## Tech Stack
- Language: Golang
- Framework: Echo
- ORM: Gorm
- IDE: Goland
- Database: MySQL
- Authentication: JWT
- ID Generator: Google UUID
- Payment Gateway: Midtrans
- VCS: Github
- Containerization: Docker
- Deployment: DigitalOcean Droplet
- Cloud Storage: DigitalOcean Spaces
- API Testing: Postman
- 
## ERD
![ERD](https://alterra.sgp1.cdn.digitaloceanspaces.com/erd.png)

## API DOCS
[SWAGGER](https://app.swaggerhub.com/apis/WAHYUUDIN2811_1/ticket-reservation/1.0.0)

## Getting Started
1. Clone repository
```
git clone https://github.com/wdinx/mini_project.git
```
2. Copy .env.example to .env
```
cp .env.example .env
```
3. Fill .env with your configuration
4. Install dependencies
```
go mod download
```
5. Run the application
```
go run main.go
```