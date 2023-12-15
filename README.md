# Go API Documentation
This project implements a basic RESTful API in Go using the Gin web framework and Gorm ORM for SQLite. The API allows you to perform CRUD operations on items.

## Docker Image
```
docker run -p 7575:8080 dhirajtoor/gocurd:1.0
```
## Endpoints
### Create Item Endpoint: POST `/create`
* Method: POST
* Request Payload: JSON Object with the following attributes: `name (string)`: The name of the item.

#### Example: 
```
{
    "name": "Dhiraj Toor"
}
```
#### Create Response:
- Status Code: 201 Created
- JSON Object with the created item.

### Read Items: GET `/read`
- Endpoint: GET /read
- Method: GET

#### Read Response:
- Status Code: 200 OK
- JSON Array containing all items.

### Update Item Endpoint: POST `/update/:id`
* Method: Endpoint: PUT /update/:id
* Method: PUT

#### URL Parameters: id (integer): 
- The ID of the item to be updated.

#### Request Payload: JSON Object with the following attributes:
- name (string): The updated name of the item..

#### Example: 
```
{
    "name": "Dhiraj"
}
```
#### Create Response:
- Status Code: 200 OK
- JSON Object with a success message.

### Read Items: GET `/delete/:id`
- Endpoint: GET /delete/:id
- Method: DELETE

#### URL Parameters: id (integer): 
- id (integer): The ID of the item to be deleted.

#### Read Response:
- Status Code: 200 OK
- JSON Object with a success message.
- 
### For test the app run below command.

```
go test
```
