Driver Location API\
Two endpoints are available\
POST: http://127.0.0.1:8080/api/drivers/save 
* It accepts application/json and multipart/form-data content types. It means you can pass single location object in request body or file as form data
* single driverLocation object example
```json
{
  "type": "Point",
  "location": {
    "longitude": 49.2174162,
    "latitude": 28.92430724
  }
}
```

* 'drivers' key and csv file


GET: http://127.0.0.1:8080/api/drivers/search
* It accepts json in request body, as given below
```json
{
  "coordinates": {
    "longitude": -73.9667,
    "latitude": 40.78
  },
  "radius": 9750000
}
```

