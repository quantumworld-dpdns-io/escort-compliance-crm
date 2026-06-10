*** Settings ***
Library    RequestsLibrary

*** Variables ***
${BASE_URL}    http://localhost:8080

*** Test Cases ***
A03 SQL Injection - Login Field
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary
    ...    email=admin'--
    ...    password=anything
    ${response}=    POST On Session    api    /api/v1/auth/login    json=${body}    headers=${headers}    expected_status=any
    Should Not Be Equal As Integers    ${response.status_code}    200

A03 SQL Injection - Search Field
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    query=' UNION SELECT * FROM users --
    ${response}=    POST On Session    api    /api/v1/companions/search    json=${body}    headers=${headers}    expected_status=any
    Should Be True    ${response.status_code} != 500

A03 SQL Injection - Booking ID
    Create Session    api    ${BASE_URL}    verify=${False}
    ${response}=    GET On Session    api    /api/v1/bookings/1' OR '1'='1    expected_status=any
    Should Be True    ${response.status_code} == 400 or ${response.status_code} == 404

A03 NoSQL Injection - MongoDB Style
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    email=${{"$ne": ""}}    password=${{"$ne": ""}}
    ${response}=    POST On Session    api    /api/v1/auth/login    json=${body}    headers=${headers}    expected_status=any
    Should Not Be Equal As Integers    ${response.status_code}    200

A03 Command Injection
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    name="; cat /etc/passwd"
    ${response}=    POST On Session    api    /api/v1/companions    json=${body}    headers=${headers}    expected_status=any
    Should Be True    ${response.status_code} != 500

A03 XSS Prevention
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    bio=<script>alert('xss')</script>
    ${response}=    POST On Session    api    /api/v1/companions    json=${body}    headers=${headers}    expected_status=any
    # Response should not contain unescaped script
    Should Not Contain    ${response.text}    <script>

A03 LDAP Injection
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    email=*)(&)    password=*)
    ${response}=    POST On Session    api    /api/v1/auth/login    json=${body}    headers=${headers}    expected_status=any
    Should Not Be Equal As Integers    ${response.status_code}    200

A03 XPath Injection
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    query=' or '1'='1
    ${response}=    POST On Session    api    /api/v1/companions/search    json=${body}    headers=${headers}    expected_status=any
    Should Be True    ${response.status_code} != 500
