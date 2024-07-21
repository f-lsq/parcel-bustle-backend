# Parcel Bustle (Backend)
[Parcel Bustle](#) is a cutting-edge courier service app designed to streamline your courier and logistic needs. With this app, you can manage parcel deliveries and track shipments in real-time with no hassle.

<p align="center">
  <img src="assets/images/readme/logo.webp#gh-light-mode-only" width="400" margin="auto">
  <img src="assets/images/readme/logo-dark.webp#gh-dark-mode-only" width="400" margin="auto">
</p>

This app is targeted for last-mile delivery services providers, specialising in transporting good from a distribution center to the end customer's doorstep. It allows workers to save parcel details into the system either using an Optical Character Recognition (OCR) tool, or by manual input. Customers will be contactable via WhatsApp with the click of a button, while the host can manage the parcels and workers' permission through the frontend interface.

* [Frontend Repository](https://github.com/f-lsq/parcel-bustle) (React Native with Redux, Expo)
* [Backend Repository](https://github.com/f-lsq/parcel-bustle-backend) (Golang - Gin, GORM, MySQL)

## Table of Contents
1. [System Design](#system-design)
2. [Functionalities](#functionalities)
3. [Technologies Used](#technologies-used)
4. [Deployment](#deployment)
5. [References](#references)

## System Design
### Entity Relationship Diagram
![Entity Relationship Diagram](assets/images/readme/erd.png#gh-light-mode-only)
![Entity Relationship Diagram](assets/images/readme/erd-dark.png#gh-dark-mode-only)
### SQL Schema Diagram
Given the relationship between entities, a relational DB (in this case MySQL) would be more ideal to store the data. The schema is as shown below.
![SQL Schema](assets/images/readme/schema.png#gh-light-mode-only)
![SQL Schema](assets/images/readme/schema-dark.png#gh-dark-mode-only)

Do note that the *password* field will be hashed using [bcrypt](https://en.wikipedia.org/wiki/Bcrypt) before it is stored in the DB. 

### System Architecture
![System Architecture](assets/images/readme/system-architecture.png#gh-light-mode-only)
![System Architecture](assets/images/readme/system-architecture-dark.png#gh-dark-mode-only)

An Object Relational Mapper (ORM) is used due to its
1. **Flexibility**: Isolates the code from the database queries, making it easier to amend the codes if the database technology were to be changed (e.g. from MySQL to PostgresSQL)
2. **Simplicity**: Performs data queries and manipulation using Golang rather than SQL language 
3. **Security**: Prevents injection attack by using the [`database/sql`](https://gorm.io/docs/security.html) argument placeholders to contruct the SQL statements

### API Routes Overview
![API Routes](assets/images/readme/api-routes.png#gh-light-mode-only)
![API Routes](assets/images/readme/api-routes-dark.png#gh-dark-mode-only)

The detailed information of each API routes can be found below.
* [Workers](#api-routes-for-workers)
* [Parcels](#api-routes-for-parcels)
* [Customers](#api-routes-for-customers)
* [Hosts](#api-routes-for-hosts)
* [Authentication](#api-routes-for-authentication)

### API Routes for Workers
Base URL for workers `http://{domain-name}.com/api/workers`.
| Endpoint | Method | Description | Request Body | Response | Authentication Response |
|----------|--------|-------------|--------------|----------|-------------------------|
| `/` | POST | Creates a new worker from request data sent as JSON | { <br>&emsp;"username": string, <br>&emsp;"email": string, <br>&emsp;"password": string,  <br>&emsp;"first_name": string, <br>&emsp;"last_name": string, <br>&emsp;"contact": string, <br>&emsp;"profile_image": string <br> } | Status 200 - returns a success message <br>Status 400 - returns an error message <br>Status 500 - returns an error message |   |
| `/` | GET | Get a list of all workers, returned as JSON |   | Status 200 - returns data of all workers <br>Status 500 - returns an error message |   |
| `/:workerId` | GET | Get a worker by their ID, returning worker data as JSON |   | Status 200 - returns data of the specific worker <br>Status 404 - returns an error message indicating that the worker is not found <br>Status 500 - returns an error message |   |
| `/:workerId` | PUT | Update data of a specific worker | { <br>&emsp;"username": string, <br>&emsp;"email": string, <br>&emsp;"password": string,  <br>&emsp;"first_name": string, <br>&emsp;"last_name": string, <br>&emsp;"contact": string, <br>&emsp;"profile_image": string <br> } | Status 200 - returns data of the updated worker <br>Status 400 - returns an error message <br>Status 404 - returns an error message indicating that the worker is not found <br>Status 500 - returns an error message | Status 401 - returns an error message requiring login <br>Status 403 - returns an error message for invalid access token |
| `/:workerId` | DELETE | Delete a specific worker |  | Status 200 - returns data of the deleted worker <br>Status 404 - returns an error message indicating that the worker is not found <br>Status 500 - returns an error message | Status 401 - returns an error message requiring login <br>Status 403 - returns an error message for invalid access token |

### API Routes for Parcels
Base URL for parcels `http://{domain-name}.com/api/parcels`.
| Endpoint | Method | Description | Request Body | Response | Authentication Response |
|----------|--------|-------------|--------------|----------|-------------------------|
| `/` | POST | Creates a new parcel from request data sent as JSON | { <br>&emsp;"status": string, <br>&emsp;"delivery_address": string, <br>&emsp;"return_address": string, <br>&emsp;"deliver_by": datetime, <br>&emsp;"delivered_image": string, <br>&emsp;"payment_type": string, <br>&emsp;"payment_mode": boolean <br> } | Status 200 - returns a success message <br>Status 400 - returns an error message <br>Status 500 - returns an error message | Status 401 - returns an error message requiring login <br>Status 403 - returns an error message for invalid access token |
| `/` | GET | Get a list of all parcels, returned as JSON |   | Status 200 - returns data of all parcels <br>Status 500 - returns an error message |   |
| `/:parcelId` | GET | Get a parcel by its ID, returning parcel data as JSON |   | Status 200 - returns data of the specific parcel <br>Status 404 - returns an error message indicating that the parcel is not found <br>Status 500 - returns an error message |   |
| `/:parcelId` | PUT | Update data of a specific parcel | { <br>&emsp;"status": string, <br>&emsp;"delivery_address": string, <br>&emsp;"return_address": string, <br>&emsp;"deliver_by": datetime, <br>&emsp;"delivered_image": string, <br>&emsp;"payment_type": string, <br>&emsp;"payment_mode": boolean <br> } | Status 200 - returns data of the updated parcel <br>Status 400 - returns an error message <br>Status 404 - returns an error message indicating that the parcel is not found <br>Status 500 - returns an error message | Status 401 - returns an error message requiring login <br>Status 403 - returns an error message for invalid access token |
| `/:parcelId` | DELETE | Delete a specific parcel |  | Status 200 - returns data of the deleted parcel <br>Status 404 - returns an error message indicating that the parcel is not found <br>Status 500 - returns an error message | Status 401 - returns an error message requiring login <br>Status 403 - returns an error message for invalid access token |
| `/worker/:workerId` | GET | Get a parcel by worker ID, returning parcel data as JSON |  | Status 200 - returns data of the specific worker <br>Status 404 - returns an error message indicating that the worker is not found <br>Status 500 - returns an error message | |
| `/customer/:customerId` | GET | Get a parcel by customer ID, returning parcel data as JSON |  | Status 200 - returns data of the specific customer <br>Status 404 - returns an error message indicating that the customer is not found <br>Status 500 - returns an error message | |

### API Routes for Customers
Base URL for customers `http://{domain-name}.com/api/customers`.
| Endpoint | Method | Description | Request Body | Response | Authentication Response |
|----------|--------|-------------|--------------|----------|-------------------------|
| `/` | POST | Creates a new customer from request data sent as JSON | { <br>&emsp;"first_name": string, <br>&emsp;"last_name": string, <br>&emsp;"contact": string <br> } | Status 200 - returns a success message <br>Status 400 - returns an error message <br>Status 500 - returns an error message |   |
| `/` | GET | Get a list of all customers, returned as JSON |   | Status 200 - returns data of all customers <br>Status 500 - returns an error message |   |
| `/:customerId` | GET | Get a customer by their ID, returning customer data as JSON |   | Status 200 - returns data of the specific customer <br>Status 404 - returns an error message indicating that the customer is not found <br>Status 500 - returns an error message |   |
| `/:customerId` | PUT | Update data of a specific customer | { <br>&emsp;"first_name": string, <br>&emsp;"last_name": string, <br>&emsp;"contact": string <br> } | Status 200 - returns data of the updated customer <br>Status 400 - returns an error message <br>Status 404 - returns an error message indicating that the customer is not found <br>Status 500 - returns an error message | Status 401 - returns an error message requiring login <br>Status 403 - returns an error message for invalid access token |
| `/:customerId` | DELETE | Delete a specific customer |  | Status 200 - returns data of the deleted customer <br>Status 404 - returns an error message indicating that the customer is not found <br>Status 500 - returns an error message | Status 401 - returns an error message requiring login <br>Status 403 - returns an error message for invalid access token |

### API Routes for Hosts
Base URL for hosts `http://{domain-name}.com/api/hosts`.
| Endpoint | Method | Description | Request Body | Response | Authentication Response |
|----------|--------|-------------|--------------|----------|-------------------------|
| `/` | POST | Creates a new host from request data sent as JSON | { <br>&emsp;"username": string, <br>&emsp;"email": string, <br>&emsp;"password": string,  <br>&emsp;"first_name": string, <br>&emsp;"last_name": string, <br> } | Status 200 - returns a success message <br>Status 400 - returns an error message <br>Status 500 - returns an error message |   |
| `/` | GET | Get a list of all hosts, returned as JSON |   | Status 200 - returns data of all hosts <br>Status 500 - returns an error message |   |
| `/:hostId` | GET | Get a host by their ID, returning host data as JSON |   | Status 200 - returns data of the specific host <br>Status 404 - returns an error message indicating that the host is not found <br>Status 500 - returns an error message |   |
| `/:hostId` | PUT | Update data of a specific host | { <br>&emsp;"username": string, <br>&emsp;"email": string, <br>&emsp;"password": string,  <br>&emsp;"first_name": string, <br>&emsp;"last_name": string, <br> } | Status 200 - returns data of the updated host <br>Status 400 - returns an error message <br>Status 404 - returns an error message indicating that the host is not found <br>Status 500 - returns an error message | Status 401 - returns an error message requiring login <br>Status 403 - returns an error message for invalid access token |
| `/:hostId` | DELETE | Delete a specific host |  | Status 200 - returns data of the deleted host <br>Status 404 - returns an error message indicating that the host is not found <br>Status 500 - returns an error message | Status 401 - returns an error message requiring login <br>Status 403 - returns an error message for invalid access token |

### API Routes for Authentication
Base URL for hosts `http://{domain-name}.com/api/auth`.
| Endpoint | Method | Description | Request Body | Response | Authentication Response |
|----------|--------|-------------|--------------|----------|-------------------------|
| `/login` | POST | Validate credentials and generate authorisation cookies | { <br>&emsp;"username": string, <br>&emsp;"email": string, <br>&emsp;"password": string <br> } | Status 200 - returns a success message <br>Status 400 - returns an error message <br>Status 401 - returns an error message for wrong credentials <br>Status 500 - returns an error message |   |
| `/refresh` | POST | Refresh expired access token with a refresh token |  | Status 200 - returns a new access token <br>Status 500 - returns an error message | Status 401 - returns an error message requiring login <br>Status 403 - returns an error message for invalid access token |
| `/refresh` | POST | Refresh expired access token with a refresh token |  | Status 204 - no response returned <br>Status 500 - returns an error message | Status 401 - returns an error message requiring login <br>Status 403 - returns an error message for invalid access token |

## Functionalities

## Technologies Used
### Frontend
* [React Native](https://reactnative.dev/)
* [Expo](https://expo.dev/)

### Backend
* [Gin](https://gin-gonic.com/docs/) - Server environment. [GORM](https://gorm.io/), [CompileDaemon](https://github.com/githubnemo/CompileDaemon), [godotenv](https://github.com/joho/godotenv), [bcrypt](golang.org/x/crypto/bcrypt), [jwt](github.com/golang-jwt/jwt), [uuid](github.com/google/uuid)
* [MySQL](https://www.mysql.com/)/[PostgresSQL](https://www.postgresql.org/) - Database management

## Deployment
### Live Links
* [React Native Frontend](#)
* [Gin Backend](#)

### Test Accounts
| Account Type | Name | Email | Password | 
|--------------|------|-------|----------|
| Worker       | [Mulan Hua](https://disney.fandom.com/wiki/Fa_Mulan) | mulanhua@parcelbustle.com    | mulanhua123@      |
| Worker       | [Shang Li](https://disney.fandom.com/wiki/Li_Shang)   | shangli@parcelbustle.com    | shangli123@      |
| Worker       | [Chi Fu](https://disney.fandom.com/wiki/Chi-Fu)   | chifu@parcelbustle.com    | shangli123@      |
| Host         | Host 1      | host1@parcelbustle.com    | *Not provided* |
| Host         | Host 2      | host2@parcelbustle.com    | *Not provided* |
| Host         | Host 3      | host3@parcelbustle.com    | *Not provided* |

## References
1. [Marcus, K. (2024 January 18). Go ORMs in 2024.](https://encore.dev/resources/go-orms)
