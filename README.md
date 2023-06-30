# Iniciador Go SDK

Welcome to the Iniciador Go SDK! This tool is made for Golang developers who want to easily integrate with our API.

If you have no idea what Iniciador is, check out our [website](https://www.iniciador.com.br/)

## 1. Description

The Iniciador SDK is a Golang library that provides a convenient way to interact with the Iniciador API.

## 2. Installation

To install the Iniciador SDK, run the following command:

```bash
go get github.com/iniciador-de-pagamentos/iniciador-sdk-go
```

## 3. Usage

To use the Iniciador SDK, import the necessary modules and create an instance of the `NewAuthClient`:

```go
  import (
    "iniciador-sdk/iniciador/auth"
  )

  func main() {
    clientID := "clientId"
    clientSecret := "clientSecret"
    environment := "dev"

    authClient := auth.NewAuthClient(clientID, clientSecret, environment)
  }
```

### 3.1 Whitelabel

#### 3.1.1 Authentication

To authenticate with the Iniciador Whitelabel, use the `AuthInterface` method:

```go
  import (
	  "iniciador-sdk/iniciador/auth"
  )

  func main() {
    authOutput, err := authClient.AuthInterface()
    if err != nil {
      fmt.Println("Authentication failed:", err)
      return
    }

    accessToken := authOutput.AccessToken
    interfaceURL := authOutput.InterfaceURL
    paymentId := authOutput.PaymentID
  }
```

- Use interfaceURL to complete the payment flow
- Use the accessToken and paymentId to verify the payment data

#### 3.1.2 Payments

To use payments services with the Iniciador Whitelabel, use the `payments` method:

##### 3.1.2.1 `get`

to get the payment details use `Get` method

```go
  import (
	  "iniciador-sdk/iniciador/payments"
  )

  func main() {
    payment, err := payments.Get(accessToken, authClient)
    if err != nil {
      fmt.Println("Get Payment failed:", err)
      return
    }
  }
```

##### 3.1.2.2 `status`

to get the payment status details use `Status` method

```go
  import (
	  "iniciador-sdk/iniciador/payments"
  )

  func main() {
    paymentStatus, err := payments.Status(accessToken, authClient)
    if err != nil {
      fmt.Println("Get Payment Status failed:", err)
      return
    }
  }
```

### 3.2 API Only

#### 3.2.1 Authentication

To authenticate with the Iniciador API, use the `Auth` method:

```go
  import (
	  "iniciador-sdk/iniciador/auth"
  )

  func main() {
    authOutput, err := authClient.Auth()
    if err != nil {
      fmt.Println("Authentication failed:", err)
      return
    }

    accessToken := authOutput.AccessToken
  }
```

#### 3.2.2 Participants

To get participants with the Iniciador API, use the `GetParticipants` method:

```go
  import (
	  "iniciador-sdk/iniciador/participants"
  )

  func main() {
    filters := &participants.ParticipantsFilter{}

    participants, err := participants.GetParticipants(accessToken, filters, authClient)
    if err != nil {
      fmt.Println("Get Participants failed:", err)
      return
    }
  }
```

#### 3.2.3 Payments

To use payments services with the Iniciador API, use the `payments` method:

##### 3.2.3.1 `send`

to send the payment use `Send` method

```go
  import (
	  "iniciador-sdk/iniciador/payments"
  )

  func main() {
    paymentPayload := &payments.PaymentInitiationPayload{
      ExternalID:    "externalId",
      ParticipantID: "c8f0bf49-4744-4933-8960-7add6e590841",
      RedirectURL:   "https://app.sandbox.inic.dev/pag-receipt",
      User: payments.User{
        Name:  "John Doe",
        TaxID: "taxId",
      },
      Amount: 133300,
      Method: "PIX_MANU_AUTO",
    }

    paymentInitiation, err := payments.Send(accessToken, paymentPayload, authClient)
    if err != nil {
      fmt.Println("Send Payments failed:", err)
      return
    }
  }
```

##### 3.2.3.2 `get`

to get the payment details use `Get` method

```go
  import (
	  "iniciador-sdk/iniciador/payments"
  )

  func main() {
    payment, err := payments.Get(accessToken, authClient)
    if err != nil {
      fmt.Println("Get Payment failed:", err)
      return
    }
  }
```

##### 3.2.3.3 `status`

to get the payment status details use `Status` method

```go
  import (
	  "iniciador-sdk/iniciador/payments"
  )

  func main() {
    paymentStatus, err := payments.Get(accessToken, authClient)
    if err != nil {
      fmt.Println("Get Payment Status failed:", err)
      return
    }
  }
```

## Help and Feedback

If you have any questions or need assistance regarding our SDK, please don't hesitate to reach out to us. Our dedicated support team is here to help you integrate with us as quickly as possible. We strive to provide prompt responses and excellent support.

We also highly appreciate any feedback you may have. Your thoughts and suggestions are valuable to us as we continuously improve our SDK and services. We welcome your input and encourage you to share your thoughts with us.

Feel free to contact us by sending an email to suporte@iniciador.com.br. We look forward to hearing from you and assisting you with your integration.
