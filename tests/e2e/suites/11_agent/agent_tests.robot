*** Settings ***
Library    RequestsLibrary
Resource   ../../resources/api_keywords.robot

*** Test Cases ***
Compliance Agent Check
    Create API Session
    ${body}=    Create Dictionary
    ...    entity_type=companion
    ...    entity_id=test-entity
    ...    jurisdiction=US-CA
    ${response}=    POST On Session    api    ${BASE_URL}/api/v1/agents/compliance/check    json=${body}    headers=${HEADERS}
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 201

Screening Agent Initiate
    Create API Session
    ${body}=    Create Dictionary
    ...    entity_id=test-entity
    ...    screening_type=background
    ${response}=    POST On Session    api    ${BASE_URL}/api/v1/agents/screening/initiate    json=${body}    headers=${HEADERS}
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 201

Scheduling Agent Optimize
    Create API Session
    ${body}=    Create Dictionary
    ...    companion_id=companion-123
    ...    date_range=${{"start": "2025-06-01", "end": "2025-06-30"}}
    ${response}=    POST On Session    api    ${BASE_URL}/api/v1/agents/scheduling/optimize    json=${body}    headers=${HEADERS}
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 201

Agent Health
    Create API Session
    ${response}=    GET On Session    api    ${BASE_URL}/api/v1/agents/health    headers=${HEADERS}
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 201
