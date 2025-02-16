# Spy Cat Agency - CRUD Application

## Project Description
This project is a RESTful API for managing a spy cat agency. It allows managing cats, missions, and targets.

## Installation and Running

1. Clone the repository:
   ```sh
   git clone https://github.com/your-repo/spy-cat-agency.git
   cd spy-cat-agency
   ```

2. Start the database using Docker:
   ```sh
   docker-compose up -d
   ```
   
   OR

   ```sh
   make docker-run
   ``` 

## API Endpoints

### Cats (Spy Cats)
- **GET** `/cats` - Retrieve a list of cats
- **POST** `/cats` - Add a new cat
- **GET** `/cats/{id}` - Retrieve information about a cat
- **PUT** `/cats/{id}` - Update cat information (salary only)
- **DELETE** `/cats/{id}` - Remove a cat

### Missions
- **GET** `/missions` - Retrieve a list of missions
- **POST** `/missions` - Create a new mission with targets
- **GET** `/missions/{id}` - Retrieve mission details
- **PUT** `/missions/{id}` - Update a mission (mark as completed)
- **DELETE** `/missions/{id}` - Delete a mission (if not assigned to a cat)

### Targets
- **PUT** `/missions/{mission_id}/targets/{target_id}` - Update notes or mark target as complete
- **DELETE** `/missions/{mission_id}/targets/{target_id}` - Delete a target (if not completed)
- **POST** `/missions/{mission_id}/targets` - Add a new target (if the mission is not completed)

## Logging
Logging is implemented via middleware to capture HTTP requests and responses.

## Validation
- Cat breed validation via [TheCatAPI](https://api.thecatapi.com/v1/breeds)
- Notes cannot be edited if the target is completed
- A mission cannot be deleted if assigned to a cat

## Testing
Use Postman (Spy Cat Agency.postman_collection.json) or cURL to test the API.

## Author
Dmytro Rozhko

