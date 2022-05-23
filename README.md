# Go-Web-API
A simple RuneScape API that contains information about locations in game.

# How to use:
1) Clone the repo to a folder on your local machine with Go installed
2) Run the project and test any of the following endpoints:

GET - http://localhost:8080/ - Home page, simply a placeholder for a home page
GET - http://localhost:8080/all - Display all default locations generated on program start
POST - http://localhost:8080/location - Create a new location through the request body
GET - http://localhost:8080/location/{name} - Display information about a single location
PUT - http://localhost:8080/location/{name} - Update information for a single location through the request body
DELETE - http://localhost:8080/location/{name} - Delete information for a single location

# Pass information through the request body using Postman using the following JSON format as an example:
{
    "Name": "Draynor Village",
    "Kingdom": "Misthalin",
    "NPCs": 22
}

# The model for a Location is as follows:
Name - string
Kingdom - string
NPCs - int
