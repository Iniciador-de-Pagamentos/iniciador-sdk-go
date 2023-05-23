package payments

import (
	"bytes"
	"fmt"
	"iniciador-sdk/iniciador/auth"
	"iniciador-sdk/iniciador/utils"
	"net/http"
)

type User struct {
	TaxID string `json:"taxId"`
	Name  string `json:"name,omitempty"`
}

type Bank struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

type OFAccountsType string

const (
	CheckingAccount OFAccountsType = "CACC"
	SalaryAccount   OFAccountsType = "SLRY"
	SavingsAccount  OFAccountsType = "SVGS"
	PrePaidAccount  OFAccountsType = "TRAN"
)

type BankAccount struct {
	TaxID       string         `json:"taxId,omitempty"`
	Name        string         `json:"name,omitempty"`
	Number      string         `json:"number"`
	AccountType OFAccountsType `json:"accountType"`
	ISPB        string         `json:"ispb"`
	Issuer      string         `json:"issuer"`
	Bank        *Bank          `json:"bank,omitempty"`
}

type PaymentInitiationStatus string

const (
	Started                        PaymentInitiationStatus = "STARTED"
	Enqueued                       PaymentInitiationStatus = "ENQUEUED"
	ConsentAwaitingAuthorization   PaymentInitiationStatus = "CONSENT_AWAITING_AUTHORIZATION"
	ConsentAuthorized              PaymentInitiationStatus = "CONSENT_AUTHORIZED"
	ConsentRejected                PaymentInitiationStatus = "CONSENT_REJECTED"
	PaymentPending                 PaymentInitiationStatus = "PAYMENT_PENDING"
	PaymentPartiallyAccepted       PaymentInitiationStatus = "PAYMENT_PARTIALLY_ACCEPTED"
	PaymentSettlementProcessing    PaymentInitiationStatus = "PAYMENT_SETTLEMENT_PROCESSING"
	PaymentSettlementDebtorAccount PaymentInitiationStatus = "PAYMENT_SETTLEMENT_DEBTOR_ACCOUNT"
	PaymentCompleted               PaymentInitiationStatus = "PAYMENT_COMPLETED"
	PaymentRejected                PaymentInitiationStatus = "PAYMENT_REJECTED"
	Canceled                       PaymentInitiationStatus = "CANCELED"
	Err                            PaymentInitiationStatus = "ERROR"
	PaymentScheduled               PaymentInitiationStatus = "PAYMENT_SCHEDULED"
)

type Provider struct {
	TradeName string `json:"tradeName"`
	Avatar    string `json:"avatar"`
	MainColor string `json:"mainColor"`
}

type Metadata map[string]interface{}

type PaymentInitiationPayload struct {
	ID                        string                   `json:"id"`
	CreatedAt                 string                   `json:"createdAt"`
	Error                     *Error                   `json:"error,omitempty"`
	Status                    *PaymentInitiationStatus `json:"status,omitempty"`
	ExternalID                string                   `json:"externalId,omitempty"`
	EndToEndID                string                   `json:"endToEndId,omitempty"`
	TransactionIdentification string                   `json:"transactionIdentification,omitempty"`
	ClientID                  string                   `json:"clientId"`
	CustomerID                string                   `json:"customerId"`
	Provider                  *Provider                `json:"provider,omitempty"`
	ConsentID                 string                   `json:"consentId,omitempty"`
	PaymentID                 string                   `json:"paymentId,omitempty"`
	ParticipantID             string                   `json:"participantId,omitempty"`
	User                      User                     `json:"user"`
	BusinessEntity            *User                    `json:"businessEntity,omitempty"`
	Method                    string                   `json:"method"`
	PixKey                    string                   `json:"pixKey,omitempty"`
	QRCode                    string                   `json:"qrCode,omitempty"`
	Amount                    float64                  `json:"amount"`
	Date                      string                   `json:"date"`
	Description               string                   `json:"description,omitempty"`
	Metadata                  Metadata                 `json:"metadata,omitempty"`
	RedirectURL               string                   `json:"redirectURL,omitempty"`
	RedirectOnErrorURL        string                   `json:"redirectOnErrorURL,omitempty"`
	IBGE                      string                   `json:"ibge,omitempty"`
	Debtor                    *BankAccount             `json:"debtor,omitempty"`
	Creditor                  *BankAccount             `json:"creditor,omitempty"`
	Fee                       float64                  `json:"fee,omitempty"`
}

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type PaymentStatusPayload struct {
	ID                        string  `json:"id"`
	Date                      string  `json:"date"`
	ConsentID                 string  `json:"consentId,omitempty"`
	CreatedAt                 string  `json:"createdAt"`
	UpdatedAt                 string  `json:"updatedAt"`
	TransactionIdentification string  `json:"transactionIdentification,omitempty"`
	EndToEndID                string  `json:"endToEndId,omitempty"`
	Amount                    float64 `json:"amount"`
	Status                    string  `json:"status"`
	Error                     *Error  `json:"error,omitempty"`
	RedirectConsentURL        string  `json:"redirectConsentURL,omitempty"`
	ExternalID                string  `json:"externalId"`
}

func Send(accessToken string, payment *PaymentInitiationPayload, authClient *auth.AuthClient) (*PaymentInitiationPayload, error) {
	payload, err := utils.MarshalWithoutEmptyFields(payment)
	if err != nil {
		return nil, err
	}

	fmt.Println(bytes.NewBuffer(payload))

	req, err := http.NewRequest("POST", authClient.Environment+"/payments", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var paymentInitiationPayload PaymentInitiationPayload
	err = utils.HandleResponse(response, &paymentInitiationPayload)
	if err != nil {
		return nil, err
	}

	return &paymentInitiationPayload, nil
}

func Get(accessToken string, authClient *auth.AuthClient) (*PaymentInitiationPayload, error) {
	paymentId, err := utils.ExtractPaymentIDFromJWTPayload(accessToken)
	if err != nil {
		fmt.Println("Something went wrong trying to get token data:", err)
		return nil, err
	}

	url := fmt.Sprintf("%s/payments/%s", authClient.Environment, paymentId)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload PaymentInitiationPayload
	err = utils.HandleResponse(resp, &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func Status(accessToken string, authClient *auth.AuthClient) (*PaymentStatusPayload, error) {
	paymentId, err := utils.ExtractPaymentIDFromJWTPayload(accessToken)
	if err != nil {
		fmt.Println("Something went wrong trying to get token data:", err)
		return nil, err
	}

	url := fmt.Sprintf("%s/payments/%s/status", authClient.Environment, paymentId)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload PaymentStatusPayload
	err = utils.HandleResponse(resp, &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
