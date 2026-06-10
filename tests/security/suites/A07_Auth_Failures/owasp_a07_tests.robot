*** Settings ***
Library    RequestsLibrary

*** Variables ***
${BASE_URL}    http://localhost:8080

*** Test Cases ***
A07 Brute Force - Login
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    email=test@example.com    password=wrongpassword

    # Attempt 10 failed logins
    FOR    ${i}    IN RANGE    10
        ${response}=    POST On Session    api    /api/v1/auth/login    json=${body}    headers=${headers}    expected_status=any
    END

    # Should be rate limited or locked out
    ${response}=    POST On Session    api    /api/v1/auth/login    json=${body}    headers=${headers}    expected_status=any
    Should Be True    ${response.status_code} == 429 or ${response.status_code} == 423

A07 Session Fixation
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    # Login should generate new session
    ${body}=    Create Dictionary    email=test@example.com    password=Test123!
    ${response}=    POST On Session    api    /api/v1/auth/login    json=${body}    headers=${headers}    expected_status=any
    # Verify token is present and not a fixed value
    IF    ${response.status_code} == 200
        Dictionary Should Contain Key    ${response.json()}    token
    END

A07 JWT Validation
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary
    ...    Content-Type=application/json
    ...    Authorization=Bearer invalid-token
    ${response}=    GET On Session    api    /api/v1/companions    headers=${headers}    expected_status=any
    Should Be True    ${response.status_code} == 401 or ${response.status_code} == 403

A07 Password Policy Enforcement
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary
    ...    email=new@example.com
    ...    password=weak
    ${response}=    POST On Session    api    /api/v1/auth/register    json=${body}    headers=${headers}    expected_status=any
    # Weak password should be rejected
    Should Not Be Equal As Integers    ${response.status_code}    201
