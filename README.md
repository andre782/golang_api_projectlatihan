Category RESTful API Documentation
Deskripsi
API ini memungkinkan Anda untuk mengelola kategori dalam sistem. Anda dapat membuat, mengupdate, mengambil, dan menghapus kategori. API ini menggunakan autentikasi berbasis API Key dengan parameter x-API-Key di header.

Versi API
Versi: 1.0
Server
API ini dapat diakses melalui URL berikut:

URL Server: http://localhost:3020/apiandre
Autentikasi
Setiap request yang dikirimkan ke API harus menyertakan header x-API-Key untuk autentikasi.

Contoh header:

bash
Copy code
x-API-Key: <API-KEY>
Endpoints
1. List All Categories
Metode HTTP: GET
Endpoint: /categories
Deskripsi: Mengambil daftar semua kategori yang ada.
Tag: Category API
Keamanan: API membutuhkan autentikasi dengan x-API-Key
Respons (200):
json
Copy code
{
  "code": 200,
  "status": "success",
  "data": [
    {
      "id": "1",
      "name": "Electronics"
    },
    {
      "id": "2",
      "name": "Books"
    }
  ]
}
2. Create New Category
Metode HTTP: POST
Endpoint: /categories
Deskripsi: Membuat kategori baru.
Tag: Category API
Keamanan: API membutuhkan autentikasi dengan x-API-Key
Request Body:
json
Copy code
{
  "name": "Home Appliances"
}
Respons (200):
json
Copy code
{
  "code": 200,
  "status": "success",
  "data": {
    "id": "3",
    "name": "Home Appliances"
  }
}
3. Get Category by ID
Metode HTTP: GET
Endpoint: /category/{categoryId}
Deskripsi: Mengambil data kategori berdasarkan ID.
Tag: Category API
Keamanan: API membutuhkan autentikasi dengan x-API-Key
Parameter:
categoryId: ID kategori yang ingin diambil.
Respons (200):
json
Copy code
{
  "code": 200,
  "status": "success",
  "data": {
    "id": "1",
    "name": "Electronics"
  }
}
4. Update Category by ID
Metode HTTP: PUT
Endpoint: /category/{categoryId}
Deskripsi: Memperbarui data kategori berdasarkan ID.
Tag: Category API
Keamanan: API membutuhkan autentikasi dengan x-API-Key
Parameter:
categoryId: ID kategori yang ingin diperbarui.
Request Body:
json
Copy code
{
  "name": "Gadgets"
}
Respons (200):
json
Copy code
{
  "code": 200,
  "status": "success",
  "data": {
    "id": "1",
    "name": "Gadgets"
  }
}
5. Delete Category by ID
Metode HTTP: DELETE
Endpoint: /category/{categoryId}
Deskripsi: Menghapus kategori berdasarkan ID.
Tag: Category API
Keamanan: API membutuhkan autentikasi dengan x-API-Key
Parameter:
categoryId: ID kategori yang ingin dihapus.
Respons (200):
json
Copy code
{
  "code": 200,
  "status": "success"
}
Komponen
Schemas
Category:

Tipe: object
Properti:
id (string): ID kategori
name (string): Nama kategori
Contoh:

json
Copy code
{
  "id": "1",
  "name": "Electronics"
}
Create/Update Category:

Tipe: object
Properti:
name (string): Nama kategori yang baru atau yang diperbarui.
Contoh:

json
Copy code
{
  "name": "Gadgets"
}
Keamanan
CategoryAuth:
Autentikasi menggunakan API Key yang disertakan di header dengan nama x-API-Key. API Key diperlukan untuk semua endpoint.
Cara Menggunakan API
Pastikan untuk memiliki API Key yang valid. API Key harus disertakan di header setiap request.
Gunakan URL server yang disediakan (http://localhost:3020/apiandre).
Gunakan metode HTTP yang sesuai untuk setiap endpoint (GET, POST, PUT, DELETE).
Kirim request sesuai dengan format JSON yang dijelaskan dalam dokumentasi di atas.
