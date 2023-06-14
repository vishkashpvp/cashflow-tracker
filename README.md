# Cashflow Tracker

**Cashflow Tracker** is a web app which lets users to track their expenses and earnings which helps them in knowing their expenditure/spendings or profits involving in different categories like food, travel, investments and etc.

Basic features of the app are:

- Track expenses by category.
- Add and manage earnings.
- Modifying/deleting existing entries.

This app doesn't need any user permissions like reading sms/mails or anything, everything is & will always be `manual`.

## Directories

This repo contains two directories `client` and `server` which are basically frontend and backend of the app respectively.

## Tech Stack

Web app uses different tech stack based on requirement.

- Frontend:
  - [Angular](https://angular.io/)
  - [Tailwindcss](https://tailwindcss.com/)
- Backend:
  - [Go](https://go.dev/)
    - [Gin](https://gin-gonic.com/)
- Database:
  - [Mongodb](https://www.mongodb.com/)

## Development server

Navigate to client directory and run `npm start` for a dev server for frontend(client) and visit `http://localhost:4200/`.

Navigate to server directory and run `go run main.go` for a dev server for backend(server) and visit `http://localhost:8080/hello`.

### Prerequisites

Node.js, Angular CLI, Go, Mongodb needed to installed to run code on the local machine.

