# Simple Authentication Template

This project provides a straightforward implementation of session management and CSRF tokens for user authentication. It serves as a template for building secure applications using Go, utilizing the standard library for most functionalities, with minimal external dependencies.

## Features

-   **User Registration**:
    -   Users can register using their email and password.
-   **Login & Logout**:
    -   Authentication is handled with session tokens, allowing users to log in and out seamlessly.
-   **Protected Routes**:
    -   Access to certain routes, such as `/show-users`, is restricted and requires valid session and CSRF tokens for access.
-   **Basic Middleware**:
    -   A simple middleware implementation is included to handle authentication without relying on external packages.
-   **Custom Router**:
    -   A basic routing mechanism is implemented to manage different endpoints without third-party libraries.
-   **Data Storage**:
    -   For demonstration purposes, user data is stored in a simple in-memory map.
-   **Validation**:
    -   Input validation is performed using the [validator](https://github.com/go-playground/validator) package (v9.31.0+incompatible).

## Usage

This template provides a foundational structure for building secure authentication mechanisms in Go applications. You can extend it to include additional features as needed.
