# Golang API Assessment

## Setting Up
1. Clone the repository
```bash
git clone https://github.com/czhi-bin/onecv-assessment.git
```

2. Edit the `config.go` file under the folder `config` to set the configuration for PostgreSQL Database Connection. This is the default configuration. Tables will be created automatically when the server runs the first time.
```
PSQL_HOST        = "localhost"
PSQL_USER        = "postgres"
PSQL_PASSWORD    = "123456"
PSQL_PORT        = 5432
PSQL_DBNAME      = "onecv-assessment"
PSQL_TEST_DBNAME = "onecv-assessment-test"
```

3. Execute the following command in terminal. The server will run on `localhost:18000` by default. This can be changed in the `config.go` file by changing the value of `LOCAL_HOST`.
```bash
go run main.go
```


## Running the unit test
1. Change the directory to `handler` folder. Execute the following command in terminal.
```bash
cd handler
go test
```

## API Documentation
1. `POST /api/register` - Register a list of student under a specifed teacher
```javascript
Header:
{
    "Content-Type": "application/json"
}
Payload:
{
    "teacher": "teacherken@gmail.com",
    "students": 
    [
        "studentjon@gmail.com",
        "studenthon@gmail.com"
    ]
}
Success response status: HTTP 204
Failure response status: 
- HTTP 400 (When the teacher or student email is invalid)
- HTTP 500 (When the server encounters an error)
```

2. `GET /api/commonstudents` - Retrieve a list of students registed to all of the given teacher
```javascript
Request example 1: 
GET /api/commonstudents?teacher=teacherken%40gmail.com
Success response body:
{
  "students" :
    [
      "commonstudent1@gmail.com", 
      "commonstudent2@gmail.com",
      "student_only_under_teacher_ken@gmail.com"
    ]
}

Request example 2: 
GET /api/commonstudents?teacher=teacherken%40gmail.com&teacher=teacherjoe%40gmail.com
Success response body:
{
  "students" :
    [
      "commonstudent1@gmail.com", 
      "commonstudent2@gmail.com",
    ]
}

Request example 3 (No registered student under one of the teacher): 
GET /api/commonstudents?teacher=nostudent%40gmail.com&teacher=teacherjoe%40gmail.com
Success response body:
{
  "students" : []
}

Success response status: HTTP 200
Failure response status: 
- HTTP 400 (When the teacher email is invalid)
- HTTP 500 (When the server encounters an error)
```

3. `POST /api/suspend` - Suspend a specified student
```javascript
Header:
{
    "Content-Type": "application/json"
}
Payload:
{
    "student": "studentmary@gmail.com"
}

Success response status: HTTP 204
Failure response status:
- HTTP 400 (When the student email is invalid / when the student is not registered)
- HTTP 500 (When the server encounters an error)
```

4. `POST /api/retrievefornotifications` - Retrieve a list of students who can receive a given notification
```javascript
Header:
{
    "Content-Type": "application/json"
}


Payload 1:
{
    "teacher":  "teacherken@gmail.com",
    "notification": "Hello students! @studentagnes@gmail.com @studentmiche@gmail.com"
}
Success response body:
{
  "recipients":
    [
      "registeredstudent@gmail.com",
      "studentagnes@gmail.com",
      "studentmiche@gmail.com"
    ]
}

Payload 2:
{
    "teacher":  "teacherken@gmail.com",
    "notification": "Hello students!"
}
Success response body:
{
  "recipients":
    [
      "registeredstudent@gmail.com"
    ]
}

Payload 3 (No registered student / All student suspended & no student mentioned in notification):
{
    "teacher":  "teacherken@gmail.com",
    "notification": "Hello students!"
}
Success response body:
{
  "recipients": []
}

Payload 4 (No registered student / All student suspended & some student mentioned in notification):
{
    "teacher":  "teacherken@gmail.com",
    "notification": "Hello students! @studentagnes@gmail.com @studentmiche@gmail.com"
}
Success response body:
{
  "recipients": 
  [
      "studentagnes@gmail.com",
      "studentmiche@gmail.com"
  ]
}

Payload 5 (suspended student mentioned in notification):
{
    "teacher":  "allstudentsuspended@gmail.com",
    "notification": "Hello students! @suspendedstudent@gmail.com"
}
Success response body:
{
  "recipients": []
}

Payload 6 (some suspended student mentioned in notification):
{
    "teacher":  "allstudentsuspended@gmail.com",
    "notification": "Hello students! @suspendedstudent@gmail.com @studentjon@gmail.com"
}
Success response body:
{
  "recipients": 
  [
    "studentjon@gmail.com"
  ]
}

Success response status: HTTP 200
Failure response status:
- HTTP 400 (When the teacher email is invalid)
- HTTP 500 (When the server encounters an error)
```

