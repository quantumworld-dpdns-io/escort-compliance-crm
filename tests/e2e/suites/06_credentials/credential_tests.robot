*** Settings ***
Library    RequestsLibrary
Resource   ../../resources/api_keywords.robot

*** Test Cases ***
Issue Credential
    Create API Session
    ${body}=    Create Dictionary
    ...    type=verification
    ...    subject=test-subject
    ...    claims=${{"age": 25, "jurisdiction": "US-CA"}}
    ${response}=    POST On Session    api    ${CREDENTIAL_URL}    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    201
    Dictionary Should Contain Key    ${response.json()}    id

Verify Credential
    Create API Session
    ${response}=    POST On Session    api    ${CREDENTIAL_URL}/test-id/verify    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200
    Should Be Equal    ${response.json()}[valid]    ${TRUE}

Selective Disclosure
    Create API Session
    ${body}=    Create Dictionary
    ...    claims=${["age", "jurisdiction"]}
    ${response}=    POST On Session    api    ${CREDENTIAL_URL}/test-id/disclose    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200
    Dictionary Should Contain Key    ${response.json()}    proof

Revoke Credential
    Create API Session
    ${response}=    POST On Session    api    ${CREDENTIAL_URL}/test-id/revoke    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200
    Should Be Equal    ${response.json()}[status]    revoked
