*** Settings ***
Library    RequestsLibrary
Resource   ../../resources/api_keywords.robot

*** Test Cases ***
Smoke Test - Health Endpoints
    Create API Session
    ${response}=    GET On Session    api    ${BASE_URL}/health
    Should Be Equal As Integers    ${response.status_code}    200

Smoke Test - Auth Flow
    Create API Session
    ${body}=    Create Dictionary    email=test@example.com    password=Test123!
    ${response}=    POST On Session    api    ${AUTH_URL}/login    json=${body}    headers=${HEADERS}    expected_status=any
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 401

Smoke Test - Companion List
    Create API Session
    ${response}=    GET On Session    api    ${COMPANION_URL}    headers=${HEADERS}    expected_status=any
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 401

Smoke Test - Quantum Health
    Create API Session
    ${response}=    GET On Session    api    http://localhost:8090/health    expected_status=any
    Should Be Equal As Integers    ${response.status_code}    200
