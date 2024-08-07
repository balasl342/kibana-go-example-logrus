# kibana-go-example-logrus
## Overview

This project provides a RESTful API for managing student records in in-memory. The API supports CRUD operations and logs requests and errors to the Elasticsearch using Logrus hooks.

## Features

- List all students
- Get details of a specific student
- Add a new student
- Update details of a specific student
- Delete a specific student
- Logs are sent to Elasticsearch

## Getting Started

### Prerequisites

- Go 1.18 or later
- Elasticsearch instance (cloud )

### Environment Variables

- Create a `.env` file in the root directory with the following content:

    ```env
    ELASTICSEARCH_URL=https://your-elasticsearch-url:443
    ELASTICSEARCH_API_KEY=your-elasticsearch-api-key

### Installation
  ```sh
    git clone https://github.com/balasl342/kibana-go-example-logrus.git
    go mod tidy
    go run main.go
  ```

### Logging

- Logs are sent to Elasticsearch using Logrus.
- Ensure that your Elasticsearch instance is accessible and properly configured in the .env file.

### Visualize in kibana

- To view the data collected by elastic:

    * Log in to your elastic cloud account.
    * Navigate to the Discover section.
    * Select your service from the list.
    * Explore the application logs that was created using this application.
