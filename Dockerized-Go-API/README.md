## Project Name : Dockerized Go Quote API.
### Description:
    This project is a small web API that renders motivational quotes using golang. It is also containerized using Docker so that it runs efficiently on any machine even if the language and other dependencies versions used to create the project are not present on the machine it is being ran.
    There are three routes in this project: the "/" which is the root that returns a simple welcome message, the "/quote" that returns random quotes in JSON format, and the "/health" that a health status message.

### How to run locally:
    To run locally, navigate into the project directory and run "go run main.go" or "go run ."

### How to run with Docker:
Running with Docker is not complex at all. Firstly, you need Docker installed. After installation, you need to start it. I use docker desktop, so to start it, run: 

    "systemctl --user start docker-desktop"
Then:

    docker build -t yourimagename (to create an image.)  
    docker images (to see your images list.)  
    docker run --name yourcontainername -p new_port:source_code_port (e.g 7000:8080) yourimagename (to create and run a container.) 
    docker ps (to see the current running container.)  

### What I learned:

    I learned a lot while working on this project. I finally learned how to use Docker in a project and I am still grinning from ear to ear. Also, it is my second time of using Go http server, my understanding is more solid, and I am so proud of myself.