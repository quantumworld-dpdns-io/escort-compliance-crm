*** Settings ***
Library    RequestsLibrary
Resource   ../../resources/api_keywords.robot

*** Test Cases ***
List Companions
    Create API Session
    ${response}=    GET On Session    api    ${COMPANION_URL}
    Should Be Equal As Integers    ${response.status_code}    200

Create Companion Profile
    Create API Session
    ${body}=    Create Dictionary
    ...    display_name=Test Companion
    ...    bio=Professional companion
    ...    jurisdiction=US-CA
    ...    rate=150.00
    ...    rate_unit=hour
    ${response}=    POST On Session    api    ${COMPANION_URL}    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    201
    Dictionary Should Contain Key    ${response.json()}    id

Search Companions
    Create API Session
    ${body}=    Create Dictionary
    ...    jurisdiction=US-CA
    ...    min_rate=100
    ...    max_rate=200
    ${response}=    POST On Session    api    ${COMPANION_URL}/search    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200

Get Companion Availability
    Create API Session
    ${response}=    GET On Session    api    ${COMPANION_URL}/test-id/availability
    Should Be Equal As Integers    ${response.status_code}    200
