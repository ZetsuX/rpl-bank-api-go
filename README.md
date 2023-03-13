# RPL Bank API

## API Documentation
Link : [Click Here!](https://documenter.getpostman.com/view/25087235/2s93CUJVem)

## Description
An API for an imaginary RPL Bank created as the 3rd assignment of admin oprec from SE Lab (RPL). This project is created using a simple Clean Architecture principle where the codes are separated into some parts which are as follows:
- `common` which is filled by utilities functions and structs used at a lot of parts in the API
- `config` which is filled by functions used to configure stuffs for the API like database connections
- `entity` which is filled with structs used in the API (mainly Database), where in this case is the database entities :
    - Nasabah (One to Many : NoTelp & Rekening)
    - NoTelp
    - Rekening
- `middleware` which is filled with functions to act as the middle layer from the web requests to the handlers
- `handler` which functions like the combination of controller, service, and repository of Clean Architecture that handle requests by the specified routes and processes them before requesting them to the database

## Tech Stack
- Golang using Gin and GORM (Back-end)
- PostgreSQL (Database)

## Assignment Instruction
Bank RPL memiliki nasabah yang dapat membuka lebih dari 1 rekening. Buatlah sistem backend untuk bank tersebut dengan fungsionalitas sebagai berikut:

Fungsi Nasabah
- Nasabah dapat mendaftar
- Nasabah dapat mengupdate informasi diri
- Nasabah dapat menghapus akun

Fungsi Rekening
- Nasabah dapat membuat rekening
- Nasabah dapat melihat semua rekening yang dia buat
- Nasabah dapat menghapus rekening

Terdapat relasi antar Nasabah dan Rekening

Untuk fungsi-fungsi itu kayak CRUD aja yaa, gaperlu authentication. Jangan lupa dengan dokumentasinya ya temen-temen!

## Features
- Supports CRUD operations for Nasabah
    - Create (C)
        - Adding New Nasabah
    - Read (R)
        - Get All Nasabah
        - Get a Nasabah by ID
        - Get All Nasabah by search query
    - Update (U)
        - Edit Nasabah
    - Delete
        - Delete Nasabah

- Supports CRUD operations for NoTelp
    - Create (C)
        - Adding New NoTelp for Nasabah
    - Read (R)
        - Get All NoTelp
        - Get All NoTelp by Nasabah ID
        - Get a NoTelp by ID
    - Update (U)
        - Edit NoTelp
    - Delete
        - Delete NoTelp

- Supports CRUD operations for Rekening
    - Create (C)
        - Adding New Rekening for Nasabah
    - Read (R)
        - Get All Rekening
        - Get All Rekening by Nasabah ID
        - Get a Rekening by ID
    - Update (U)
        - Edit Rekening
    - Delete
        - Delete Rekening

## Hardships I Felt
- Me, who never used Gin and GORM for Golang before of course have to learn more about how to use them, especially the functions that are available and needed to be utilised in creating the API
- The architecture used was kind of different with the one I created by myself on my 2nd assignent, therefore there was some readjusting needed for me

## Things I Learned
- By doing this assignment, I've started knowing more about Gin and GORM usage in programming using Golang. 
- I've also understood more about the capability and functionality of Gin and GORM to further make our works as a Golang Backend Programmer easier rather than using native Golang as the one used in the 2nd assignment.
- Beside those, I've also learned more about API and it's aspect to the point where I started considering about the usage and always trying to improve the efficiency of it in every task.
