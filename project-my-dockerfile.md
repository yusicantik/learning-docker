1. Create a new Golang project that serve http
- Path "/", do print string/html text of your favorite wise word
- When you access path "/", also print to log/console using fmt.Println with text "Ouch!" Start HTTP on port 77
<img width="656" alt="image" src="https://github.com/yusicantik/learning-docker/assets/76909692/9262d3ed-4d1e-4fc1-b7cb-77a441f9e80e">



2. Create file "AUTHORS.md" and "LINKS.md" within same folder with "Dockerfile
- Edit "AUTHORS.md" file and fill it with your first name or your fake name
<img width="652" alt="image" src="https://github.com/yusicantik/learning-docker/assets/76909692/22bbed93-81b3-4f38-9b4a-c8e1a5511a1f">


- Edit "LINKS.md" file and fill it with your github profile link
<img width="652" alt="image" src="https://github.com/yusicantik/learning-docker/assets/76909692/b4b83d86-5481-4598-842c-21412153276d">



3. Create Dockerfile to build and run your golang project
- Use Base Image "golang:1.21"
- Set WORKDIR with /myapp
- RUN some command "go version" after WORKDIR
- COPY "AUTHORS.md" to image BEFORE run build golang (go build) COPY "LINKS.md" to image BEFORE run build golang (go build)
- Make golang build output with name "my-go-app" and make sure it will run correctly
<img width="655" alt="image" src="https://github.com/yusicantik/learning-docker/assets/76909692/4c0f608e-0e6c-4b7d-adb2-4f81a2064a06">



4. Build Dockerfile image with name "my-go-app:v2" docker run --name go-app -p 5555:77 my-go-app:v2
<img width="653" alt="image" src="https://github.com/yusicantik/learning-docker/assets/76909692/6fee914f-fd6b-4f5c-9ec1-f049f01d09c7">



6. Run the image into container with name "go-app" and expose to port host 5555 docker build -t my-go-app:v2
<img width="654" alt="image" src="https://github.com/yusicantik/learning-docker/assets/76909692/605a827e-7a04-47fa-964b-63c119305967">
<img width="960" alt="image" src="https://github.com/yusicantik/learning-docker/assets/76909692/472765a1-1b4c-408d-9869-3ea1f88f8c59">



7. Access your golang inside docker via localhost:5555 and see logs of your container "go-app"
<img width="960" alt="image" src="https://github.com/yusicantik/learning-docker/assets/76909692/6c72b780-ee44-412d-8b11-ce230f7496ef">


