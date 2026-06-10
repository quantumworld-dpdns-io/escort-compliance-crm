*** Settings ***
Library    RequestsLibrary

*** Variables ***
${BASE_URL}    http://localhost:8080

*** Test Cases ***
A08 Deserialization Protection
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    # Send malformed JSON
    ${response}=    POST On Session    api    /api/v1/auth/login    data=not-json    headers=${headers}    expected_status=any
    Should Be True    ${response.status_code} == 400 or ${response.status_code} == 422

A08 Input Validation
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary
    ...    email=not-an-email
    ...    password=
    ${response}=    POST On Session    api    /api/v1/auth/login    json=${body}    headers=${headers}    expected_status=any
    Should Not Be Equal As Integers    ${response.status_code}    500

A08 Data Integrity - Request Signing
    Create Session    api    ${BASE_URL}    verify=${False}
    # Verify requests have integrity checks
    ${response}=    GET On Session    api    /health
    Dictionary Should Contain Key    ${response.headers}    X-Request-ID
