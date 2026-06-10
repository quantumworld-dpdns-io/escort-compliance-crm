*** Settings ***
Library    RequestsLibrary
Resource   ../../resources/api_keywords.robot

*** Test Cases ***
Create Booking
    Create API Session
    ${body}=    Create Dictionary
    ...    companion_id=companion-123
    ...    start_time=2025-06-15T14:00:00Z
    ...    end_time=2025-06-15T16:00:00Z
    ...    jurisdiction=US-CA
    ...    location=Los Angeles, CA
    ${response}=    POST On Session    api    ${BOOKING_URL}    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    201
    Dictionary Should Contain Key    ${response.json()}    id
    Should Be Equal    ${response.json()}[status]    pending

Confirm Booking
    Create API Session
    ${response}=    POST On Session    api    ${BOOKING_URL}/test-id/confirm    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200
    Should Be Equal    ${response.json()}[status]    confirmed

Cancel Booking
    Create API Session
    ${response}=    POST On Session    api    ${BOOKING_URL}/test-id/cancel    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200
    Should Be Equal    ${response.json()}[status]    cancelled

Complete Booking
    Create API Session
    ${response}=    POST On Session    api    ${BOOKING_URL}/test-id/complete    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200
    Should Be Equal    ${response.json()}[status]    completed
