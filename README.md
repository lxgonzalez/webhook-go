# Webhook using go

## Tech Stack

![Go_Icon](https://img.icons8.com/?size=49&id=44442&format=png&color=000000)
## Deployment

To run the API locally, follow these steps:

### 1. Clone the Repository
Clone this repository to your local machine:
```bash
git https://github.com/lxgonzalez/webhook-go
cd webhook-go
```
### 2. Run the Application
```bash
go run main.go
```
### 3. Testing 
- Open postman or insomnia
- Create a new request
- Set the request method to POST
- Set the request URL http://localhost:8080/webhook
- In the Body tab, choose raw and set the format to JSON
- Add the following JSON payload to the body:
```bash
{
    "event": "example_event",
    "message": "Hello, webhook!"
}
```
### DEMO
![image](https://github.com/user-attachments/assets/944b9221-35e8-4d47-872e-5e9a641cae27)

## Authors

- [@lxgonzalez](https://github.com/lxgonzalez)
