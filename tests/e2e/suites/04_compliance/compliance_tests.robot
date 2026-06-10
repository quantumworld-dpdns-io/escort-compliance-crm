*** Settings ***
Library    RequestsLibrary
Resource   ../../resources/api_keywords.robot

*** Test Cases ***
List Jurisdictions
    Create API Session
    ${response}=    GET On Session    api    ${COMPLIANCE_URL}/jurisdictions
    Should Be Equal As Integers    ${response.status_code}    200

Compliance Check
    Create API Session
    ${body}=    Create Dictionary
    ...    jurisdiction_id=US-CA
    ...    entity_type=companion
    ...    entity_id=test-entity
    ${response}=    POST On Session    api    ${COMPLIANCE_URL}/check    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200
    Dictionary Should Contain Key    ${response.json()}    compliant
    Dictionary Should Contain Key    ${response.json()}    score

Get Compliance Alerts
    Create API Session
    ${response}=    GET On Session    api    ${COMPLIANCE_URL}/alerts
    Should Be Equal As Integers    ${response.status_code}    200
