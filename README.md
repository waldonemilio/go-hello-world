# Hello World Extravaganza

A simple yet elegant web application showcasing the power of Dockerized Go for dynamic content generation.

## Getting Started

**Prerequisites**

* Docker and Docker Compose installed on your system

**Instructions**

1. Clone this repository:

   ```bash
   git clone https://github.com/waldonemilio/go-hello-world.git
   ```

2. Navigate to the project directory:

    ```bash
    cd hello-world-extravaganza
    ```

3. Build and run the Docker image:

    ```bash
    docker compose up -d --build
    ```

4. Access the application in your web browser at: `http://localhost:3000`

## Project Structure

* `main.go`: Contains the Go application logic, number generation, and template handling. Sets up routes to serve static files.
* `index.html`:  Defines the layout with sections for navigation, the 'hero' element, and dynamic number display. Uses Go templating features.
* `style.css`:  Handles basic styling and organization for layout elements.
* `Dockerfile`: Specifies instructions for building the Docker image of the application.
* `docker-compose.yml`: Automates building and running the container.

## Customization

* **Go Logic (`main.go`)**: Extend the number generation, or create entirely new dynamic content powered by Go.
* **Templating (`index.html`)**: Add new sections, data elements, and change the integration of Go's templating engine.
* **Styling (`style.css`)**: Expand upon the stylesheet to fine-tune the aesthetic of your application.  You could adopt a full CSS framework such as Bootstrap.

## Future Development

* **Interactivity with JavaScript (`script.js`)**:  Make menus interactive, create input forms, or add more dynamic behaviors with JavaScript code.
* **Multi-Page Application**: Create multiple routes and HTML templates for different pages within the site.
* **Database Integration**:  For adding functionality with real storage such as logins, product tracking, or other features that store information for subsequent uses.
