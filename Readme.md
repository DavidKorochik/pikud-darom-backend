# Pikud Darom Backend Service

This is the backend service for the project I'm developing for Pikud Darom

## How to start the program?

```bash
go run main.go
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate

## Database Diagram

![Database Diagram](/images/database-diagram.jpg)

## Models Schema

# The Issue Model

```golang
	IssueID               uuid.UUID
	Date                  string
	Hour                  string
	Unit                  string
	Topic                 string
	SpecificTopic         string
	MonitoringType        string
	UserID                uuid.UUID
	MonitoringSystem      string
	IssueCause            string
	ResponsibleDepartment string
	Status                string
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             time.Time
```

# The User Model

```golang
	UserID         uuid.UUID
	FirstName      string
	LastName       string
	ArmyEmail      string
	PersonalNumber string
	Department     string
	Issues         *[]Issue
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
```
