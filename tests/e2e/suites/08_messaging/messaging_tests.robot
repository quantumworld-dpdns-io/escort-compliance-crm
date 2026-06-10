*** Settings ***
Library    RequestsLibrary
Resource   ../../resources/api_keywords.robot

*** Test Cases ***
Send Message
    Create API Session
    ${body}=    Create Dictionary
    ...    recipient_id=user-456
    ...    content=Hello, this is a test message
    ${response}=    POST On Session    api    ${BASE_URL}/api/v1/messages    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    201
    Dictionary Should Contain Key    ${response.json()}    id

List Messages
    Create API Session
    ${response}=    GET On Session    api    ${BASE_URL}/api/v1/messages    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200

Mark Message Read
    Create API Session
    ${response}=    POST On Session    api    ${BASE_URL}/api/v1/messages/test-id/read    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200

Get Conversation
    Create API Session
    ${response}=    GET On Session    api    ${BASE_URL}/api/v1/conversations/test-id    headers=${HEADERS}
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 404
