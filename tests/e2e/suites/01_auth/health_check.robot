*** Settings ***
Library    RequestsLibrary
Resource   ../resources/api_keywords.robot

*** Test Cases ***
Health Check
    Create API Session
    ${response}=    GET On Session    api    ${BASE_URL}/health
    Should Be Equal As Integers    ${response.status_code}    200
    Dictionary Should Contain Key    ${response.json()}    status
    Should Be Equal    ${response.json()}[status]    healthy

Gateway Health
    Create API Session
    ${response}=    GET On Session    api    ${API_URL}/../../health
    Should Be Equal As Integers    ${response.status_code}    200

Auth Service Health
    Create API Session
    ${response}=    GET On Session    api    http://localhost:8081/health
    Should Be Equal As Integers    ${response.status_code}    200
    Should Be Equal    ${response.json()}[service]    auth

Companion Service Health
    Create API Session
    ${response}=    GET On Session    api    http://localhost:8082/health
    Should Be Equal As Integers    ${response.status_code}    200

Booking Service Health
    Create API Session
    ${response}=    GET On Session    api    http://localhost:8083/health
    Should Be Equal As Integers    ${response.status_code}    200

Compliance Service Health
    Create API Session
    ${response}=    GET On Session    api    http://localhost:8084/health
    Should Be Equal As Integers    ${response.status_code}    200

Screening Service Health
    Create API Session
    ${response}=    GET On Session    api    http://localhost:8085/health
    Should Be Equal As Integers    ${response.status_code}    200

Credentials Service Health
    Create API Session
    ${response}=    GET On Session    api    http://localhost:8086/health
    Should Be Equal As Integers    ${response.status_code}    200

Quantum Service Health
    Create API Session
    ${response}=    GET On Session    api    http://localhost:8090/health
    Should Be Equal As Integers    ${response.status_code}    200
