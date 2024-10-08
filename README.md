# Profile Generator

This profile generator provides a REST API to serve profile information in JSON format.

## How to Build

1. **Clone the Repository:**

   ```sh
   git clone git@github.com:MurmurationsNetwork/profile-generator.git
   cd profile-generator
   ```

2. **Set Up Docker Buildx:**

   Ensure you have Docker Buildx installed and configured. Docker Buildx is a Docker CLI plugin that extends the Docker command with the full support of the features provided by the Moby BuildKit builder toolkit.

   ```sh
   docker buildx create --use
   docker buildx inspect --bootstrap
   ```

3. **Build a Multi-Platform Docker Image:**

   Use Docker Buildx to build and push a multi-platform Docker image that supports both ARM64 and AMD64.

   ```sh
   docker buildx build --platform linux/amd64,linux/arm64 -t <your-dockerhub-username>/profile-generator --push .
   ```

## Run Inside Your Server

1. **Login to Your Server:**

   ```sh
   ssh -i <path-to-your-private-key> root@<your-server-ip>
   ```

2. **Pull the Docker Image from Docker Hub:**

   ```sh
   docker pull <your-dockerhub-username>/profile-generator
   ```

3. **Run the Docker Container:**

   ```sh
   docker stop profile-generator
   docker rm -f profile-generator
   docker run -d -p 80:8080 --name profile-generator <your-dockerhub-username>/profile-generator
   ```

4. **Verify the Deployment:**

   Visit `http://<your-server-ip>/profile/small/{any-number}` in a web browser or use `curl` to verify that the application is running:

   ```sh
   curl http://<your-server-ip>/profile/small/120
   ```

   You should receive a JSON response similar to:

   ```json
   {
     "linked_schemas": [
       "test_schema-v2.0.0"
     ],
     "name": "IC3 Dev",
     "latitude": 11.111111,
     "longitude": 12.121212
   }
   ```
