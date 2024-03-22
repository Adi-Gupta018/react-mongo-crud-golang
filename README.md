# React-Mongo-Crud-Golang

This project is a full-stack web application built with Golang and React. The backend utilizes Golang and the Gin framework for a RESTful API that interacts with a MongoDB database. The frontend is built with React using Create React App, providing a user interface for CRUD (Create, Read, Update, Delete) operations on citizen data.
![image](https://github.com/Adi-Gupta018/react-mongo-crud-golang/assets/94818088/f8fe2c15-afe2-4085-b09d-9c5e3d9aef22)


## Tech Stack

Backend:
```
  Golang,

  Gin (web framework)

  MongoDB (database)

  godotenv (for environment variables)
```
Frontend:
```
  React (built with create-react-app)
```

## Running the project


Prerequisites:

1. Make sure you have Golang and Git installed on your system.

 2. Install MongoDB and set it up according to their documentation.

Clone the repository:

    git clone https://github.com/Adi-Gupta018/react-mongo-crud-golang.git

## Set up Backend:

  Install dependencies:
```Bash
cd <your-repo-name>
go mod download
```

## Set up Frontend:

   Navigate to the frontend directory:
```Bash   
cd frontend
```
Install dependencies:
```Bash

npm install or yarn install
```
Start the development server:
```Bash
npm start or yarn start
```

API Endpoints:

    GET /citizens/:id: Retrieves a citizen by ID.
    GET /citizens: Retrieves all citizens.
    POST /citizens: Creates a new citizen.
    PUT /citizens/:id: Updates a citizen by ID.
    DELETE /citizens/:id: Deletes a citizen by ID.

Note:

    CORS is enabled using the gin-cors middleware to allow requests from any origin in the backend.
    Refer to the code for detailed implementation of each endpoint and data model used.


## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.


## License

[MIT](https://choosealicense.com/licenses/mit/)
