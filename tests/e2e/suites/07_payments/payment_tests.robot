*** Settings ***
Library    RequestsLibrary
Resource   ../../resources/api_keywords.robot

*** Test Cases ***
Create Payment
    Create API Session
    ${body}=    Create Dictionary
    ...    booking_id=booking-123
    ...    amount=300.00
    ...    currency=USD
    ...    method=stripe
    ${response}=    POST On Session    api    ${BASE_URL}/api/v1/payments    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    201
    Dictionary Should Contain Key    ${response.json()}    id

Process Payment
    Create API Session
    ${response}=    POST On Session    api    ${BASE_URL}/api/v1/payments/test-id/process    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200

Refund Payment
    Create API Session
    ${body}=    Create Dictionary    reason=customer_request
    ${response}=    POST On Session    api    ${BASE_URL}/api/v1/payments/test-id/refund    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200

List Payments
    Create API Session
    ${response}=    GET On Session    api    ${BASE_URL}/api/v1/payments    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200

Escrow Hold
    Create API Session
    ${body}=    Create Dictionary
    ...    booking_id=booking-456
    ...    amount=500.00
    ${response}=    POST On Session    api    ${BASE_URL}/api/v1/payments/escrow    json=${body}    headers=${HEADERS}
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 201
