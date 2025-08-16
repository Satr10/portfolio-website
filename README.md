# Portfolio Website

This is a portfolio website with a terminal-like interface, built with Go, Fiber, and HTMX.

## About

This project is a personal portfolio website designed to showcase my projects and skills in a unique and interactive way. The terminal-style interface allows users to navigate the site using commands, similar to a command-line application.

## Features

The website supports the following commands:

-   `welcome`: Displays the welcome message.
-   `help`: Displays a list of available commands.
-   `about`: Shows information about me.
-   `projects`: Lists my projects.
-   `contact`: Provides my contact information.
-   `clear`: Clears the terminal screen.

## Technologies Used

-   **Backend:** Go, Fiber
-   **Frontend:** HTMX, TailwindCSS
-   **Templating:** Django-like template engine for Go
-   **Live Reloading:** Air

## Installation

To run this project locally, you will need to have Go and Node.js installed.

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/Satr10/portfolio-website.git
    cd portfolio-website
    ```

2.  **Install Go dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Install frontend dependencies:**
    ```bash
    npm install
    ```

4.  **Build the CSS:**
    ```bash
    npm run build:css
    ```

5.  **Run the application:**
    You can run the application using `go run`:
    ```bash
    go run main.go
    ```
    The application will be available at `http://localhost:5001`.

    For development, you can use the `air` tool for live reloading:
    ```bash
    air
    ```

## Usage

Once the application is running, you can open it in your browser. You will be greeted with a terminal-like interface. Type `help` to see a list of available commands and start exploring the portfolio.
