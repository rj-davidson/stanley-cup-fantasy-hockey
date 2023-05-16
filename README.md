# stanley-cup-fantasy-hockey

## Description
This is a fantasy hockey app that allows users to create a team and compete against other users in a league. The app will allow users to draft players, make trades, and view stats. The app will also allow users to view the current standings and view the current playoff picture.

## Website Demo
![hockey.rjd.app](https://hockey.rjd.app/)

## Table of Contents
* [Installation](#installation)
* [Usage](#usage)
* [License](#license)
* [Technologies](#technologies)

## Installation
The app is divided into two services, server and client. The database configuration is set up for PostgreSQL. The server is built with Go and has been tested on Go 1.18+. The client is built with Next.js and has been tested on Node 16.13.0+.

To install the server, navigate to the server directory and build main.go. The server requires a .env (server/.env) file with the following variables:
```
DB_HOST="<your host here>"
DB_PORT="5432"
DB_USER="<your db username>"
DB_PASS="<your password>"
DB_NAME="<your db name>"
IP_HOST="<local host or your server ip>"
DOMAIN="<your domain [optional]>"
```

To install the client, navigate to the ui directory and run the following commands:
```
npm install yarn
yarn install
yarn run dev
```
The client requires a .env (ui/.env) file with the following variables:
```
NEXT_PUBLIC_API_URL="<your server ip or domain>"
```

## Usage
The app is currently in development and is not ready for production. You're welcome to use the app for testing purposes but please do not use it for production.

## License
This project is licensed under the MIT license. See `LICENSE` for details.

## Technologies
* Go -> Ent (ORM), Fiber (Web Framework), Cron (Scheduler), Viper (Configuration)
* Next.js -> React, TypeScript, Tailwind CSS, MUI
* Database -> PostgreSQL
* Other technologies used: Docker, Nginx, Certbot, Let's Encrypt and code assistance from GitHub Copilot and ChatGPT