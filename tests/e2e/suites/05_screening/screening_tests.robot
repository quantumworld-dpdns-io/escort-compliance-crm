*** Settings ***
Library    RequestsLibrary
Resource   ../../resources/api_keywords.robot

*** Test Cases ***
Create Screening
    Create API Session
    ${body}=    Create Dictionary
    ...    entity_id=test-entity
    ...    type=background_check
    ${response}=    POST On Session    api    ${SCREENING_URL}    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    201
    Dictionary Should Contain Key    ${response.json()}    id

Get Screening Status
    Create API Session
    ${response}=    GET On Session    api    ${SCREENING_URL}/test-id/status
    Should Be Equal As Integers    ${response.status_code}    200
    Dictionary Should Contain Key    ${response.json()}    status
