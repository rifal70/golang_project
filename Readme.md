# Nama Proyek

Deskripsi singkat proyek.

## Menjalankan Aplikasi

Pastikan telah menginstal Golang di mesin. Jika belum, unduh dan instal dari [situs resmi Golang](https://golang.org/dl/).

1. Clone repositori ini ke mesin lokal:
    ```bash
    git clone https://github.com/rifal70/golang_project
    ```

2. Pindah ke direktori proyek:
    ```bash
    cd <nama-direktori-proyek>/golang_project
    ```

3. Jalankan aplikasi dengan perintah:
    ```bash
    go run main.go
    ```
    Ini akan menjalankan server HTTP lokal pada `localhost:8069`.

## Menggunakan Postman untuk Menguji Endpoint

Pastikan aplikasi Golang berjalan.

- **POST** untuk menambahkan data hewan baru:
  - URL: `http://localhost:8069/v1/animal`
  - Method: POST
  - Body (raw JSON):
    ```json
    {
      "id": 2,
      "name": "tiger",
      "class": "mammal",
      "legs": 4
    }
    ```

- **PUT** untuk memperbarui atau membuat data hewan baru:
  - URL: `http://localhost:8069/v1/animal/{id}`
  - Method: PUT
  - Body (raw JSON):
    ```json
    {
      "name": "tiger-updated",
      "class": "mammal",
      "legs": 4
    }
    ```

- **DELETE** untuk menghapus data hewan:
  - URL: `http://localhost:8069/v1/animal/{id}`
  - Method: DELETE

- **GET** untuk mendapatkan daftar semua data hewan:
  - URL: `http://localhost:8069/v1/animal`
  - Method: GET

- **GET** untuk mendapatkan hewan berdasarkan ID:
  - URL: `http://localhost:8069/v1/animal/{id}`
  - Method: GET

Pastikan untuk mengganti `{id}` dengan ID yang sesuai saat menggunakan permintaan PUT, DELETE, atau GET berdasarkan ID.