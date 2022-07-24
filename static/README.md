# cnqr
This API will be hosted on Cloud Run, a serverless platform for running containers. API routes will be documented here in the README and commented in the code. This README tracks along with an Open API Spec created in Stoplight Studio. 

## Authentication/Authorization
Unless otherwise specified, all endpoints require a valid firebase identity token to be passed in the `Authorization` header in the format `Bearer {token}`. This token is validated server side to make sure it has not expired or been tampered with. The client side SDK should make this token easy to retrieve and use.

If a request is made without authentication, the API will return a 401. If a request is made with authentication, but the user does not have proper authorization to access the resource, the API will return a 403.

Firebase Auth allows for custom claims to be embedded in identity tokens returned from the service. This API makes use of those custom claims to determine if a user is allowed to access a resource. So far, custom claims are:

`companyId` to limit access to only the company specified

`merchantId` to limit access to only the company specified in the merchant portal (unimplemented)

<!-- `codatCompanyId` to limit access to only the company specified -->

`admin` to limit access to the admin portal

## Routes
So far, all API paths are subject to change, according to the frontend team's needs. All returned data will be in the JSON format, and the API expects the same of all incoming data. All route information can be found in routes.go. 

### Paged Endpoints

In the response body there may be a next_page_token, if there is then it can be used to fetch the next page of results. When there is not a next page token, there are no more pages of results to fetch. There are also query parameters required for most of the endpoints. All timestamps must be in RFC3339 format. If using a page token, the other query parameters must be the same for each page requested.

### SendGrid

There are some events during which an email is sent to the user. 

- When an admin marks an application as accepted
- (plaid item error?)


### Webhooks and Events

- Plaid
    - ERROR
    - NEW_ACCOUNTS_AVAILABLE
    - PENDING_EXPIRATION
    - USER_PERMISSION_REVOKED
    - WEBHOOK_UPDATE_ACKNOWLEDGED
- Circle
    - Payouts
    - Payments
    - ACH
- Codat
    - NewCompanySynchronized
    - DataSyncError
    - DatasetDataChanged
    - DataSyncCompleted
    - DataConnectionStatusChanged

### Onboarding Process

Process begins with a call to /onboarding/start which creates the onboarding progress resource

This is all pending the Persona Integration, most likely would start with Persona, then Codat Link, then all misc info, application status, transfer info and finally finish

- Email Verification (maybe should be condensed into a requirement before starting)
    - /onboarding/confirmEmailVerification
- Collect Referral Code
    - /onboarding/collectReferralCode
    - /onboarding/joinWaitlist
- On Waitlist
- Get Started
    - /onboarding/collectStartingInfo
- Connect Bank Account
    - /onboarding/confirmConnectedBankAccount
- Connect Sales Account
    - /onboarding/corfirmConnectedSalesAccount
- Choose Service
    - /onboarding/chooseContinuousService
    - /onboarding/chooseOneTimeService
- Confirm Persona
- Status Pending
- Status Rejected
- Status Accepted
    - /onboarding/confirmAccepted
- Collect Transfer Info
    - /onboarding/collectTransferInformation
- Rates and Agreements
    - /onboarding/confirmRatesAndAgreements
- Finished

Once an onboarding progress is Finished, the user will be redirected to the merchant portal whenever the redirect handler is called


## ALL TODOS
Persona Integration (waiting on document information)
Sendgrid Integration (client is ready, add in mail sending when template is done)
Data lifecycle (transaction logic) - need for few onboarding steps (waiting for Persona)
Commercial API auth and management

Verify prequalification logic
Verify commercial onboarding logic, how users transfer

Do you have expected commercial users? How many and what kind of auth do you want? Will the users already have an account created?

What kind of logic or calculations are needed over the data? Some of the fields may either be redundant or need to be verified anyway, like total amounts. 

Verify schemas of requests

esig and docusign

financing offer, request, receive 