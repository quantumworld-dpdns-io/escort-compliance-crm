*** Settings ***
Library    RequestsLibrary
Library    Collections
Library    OperatingSystem
Library    String
Library    DateTime

*** Variables ***
${BASE_URL}         http://localhost:8080
${API_URL}          ${BASE_URL}/api/v1
${AUTH_URL}         ${API_URL}/auth
${COMPANION_URL}    ${API_URL}/companions
${BOOKING_URL}      ${BASE_URL}/api/v1/bookings
${COMPLIANCE_URL}   ${API_URL}/compliance
${SCREENING_URL}    ${API_URL}/screening
${CREDENTIAL_URL}   ${API_URL}/credentials
${QUANTUM_URL}      ${API_URL}/quantum

*** Keywords ***
Create API Session
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    Set Suite Variable    ${HEADERS}    ${headers}

Get Auth Token
    ${body}=    Create Dictionary    email=test@example.com    password=Test123!
    ${response}=    POST On Session    api    ${AUTH_URL}/login    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200
    ${token}=    Get From Dictionary    ${response.json()}    token
    Set Suite Variable    ${AUTH_TOKEN}    ${token}
    Set To Dictionary    ${HEADERS}    Authorization=Bearer ${AUTH_TOKEN}

Generate Test ID
    ${id}=    Evaluate    str(__import__('uuid').uuid4())
    [Return]    ${id}
