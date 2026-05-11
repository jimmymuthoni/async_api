### async API
An Asynchronous API (or Async API) is an interface that allows a client to send a request and then continue with other tasks without waiting for an immediate response.Unlike synchronous APIs (like typical REST), which block execution until the server replies, async APIs are non-blocking.

This project provides an asynchronous report generation system where clients submit report requests and receive a report ID immediately instead of waiting for processing to finish.

The architecture is designed to:

Handle heavy or long-running workloads.
Improve API responsiveness.
Decouple processing from request handling.
Scale workers independently.
Store generated reports reliably in cloud storage.

### Architecture:

![Alt text](/images/architecture.png)

The system consists of:

Client
API Server
PostgreSQL Database
Amazon SQS Queue
Worker Service
Amazon S3 Storage
External LoZ API

### system flow:

```

Client
   |
   | POST /reports
   v
API Server
   |
   |-- Authenticate User
   |-- Create report job in PostgreSQL
   |-- Generate report_id
   |-- Push report_id to SQS Queue
   |
   v
Immediate Response to Client
{
  "report_id": "rep_12345",
  "status": "queued"
}

                ┌──────────────────────┐
                │   reports_sqs_queue  │
                └──────────────────────┘
                           |
                           v
                    SQS Worker Service
                           |
            ┌──────────────┼──────────────┐
            |                              |
            v                              v
     Fetch Job Details              Call LoZ API
       from PostgreSQL
            |
            v
     Generate Report
            |
            v
     Upload Report to S3
            |
            v
   Update Job Status in PostgreSQL
            |
            v
Client Requests Report Status
   |
   | GET /reports/{report_id}
   v
API Server
   |
   |-- Fetch status from PostgreSQL
   |
   v
Return Response

Possible Responses:

{
  "report_id": "rep_12345",
  "status": "processing"
}

or

{
  "report_id": "rep_12345",
  "status": "completed",
  "download_url": "https://s3.amazonaws.com/api-reports/report.pdf"
}

```

### Technologies:
Backend 
- Go / Golang
- REST API
Infrastructure
- Amazon SQS
- Amazon S3
- PostgreSQL
Authentication
- JWT Authentication
- Refresh Tokens