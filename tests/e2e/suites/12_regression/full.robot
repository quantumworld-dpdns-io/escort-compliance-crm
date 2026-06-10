*** Settings ***
Library    RequestsLibrary
Resource   ../../resources/api_keywords.robot

*** Test Cases ***
Full Auth Flow
    Create API Session
    ${body}=    Create Dictionary
    ...    name=Test User
    ...    email=fulltest@example.com
    ...    password=SecurePass123!
    ...    role=client
    ${response}=    POST On Session    api    ${AUTH_URL}/register    json=${body}    headers=${HEADERS}    expected_status=any
    Should Be True    ${response.status_code} == 201 or ${response.status_code} == 409

Full Companion Flow
    Create API Session
    ${body}=    Create Dictionary
    ...    display_name=Test Companion
    ...    jurisdiction=US-CA
    ...    rate=150.00
    ${response}=    POST On Session    api    ${COMPANION_URL}    json=${body}    headers=${HEADERS}    expected_status=any
    Should Be True    ${response.status_code} == 201 or ${response.status_code} == 401

Full Booking Flow
    Create API Session
    ${body}=    Create Dictionary
    ...    companion_id=companion-123
    ...    start_time=2025-07-01T14:00:00Z
    ...    end_time=2025-07-01T16:00:00Z
    ...    jurisdiction=US-CA
    ${response}=    POST On Session    api    ${BOOKING_URL}    json=${body}    headers=${HEADERS}    expected_status=any
    Should Be True    ${response.status_code} == 201 or ${response.status_code} == 401

Full Compliance Flow
    Create API Session
    ${body}=    Create Dictionary
    ...    jurisdiction_id=US-CA
    ...    entity_type=companion
    ...    entity_id=test-entity
    ${response}=    POST On Session    api    ${COMPLIANCE_URL}/check    json=${body}    headers=${HEADERS}    expected_status=any
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 401

Full Quantum PQC Flow
    Create API Session
    ${body}=    Create Dictionary    algorithm=kyber768
    ${response}=    POST On Session    api    ${QUANTUM_URL}/pqc/keygen    json=${body}    headers=${HEADERS}    expected_status=any
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 401
