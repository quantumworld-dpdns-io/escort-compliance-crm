*** Settings ***
Library    RequestsLibrary

*** Variables ***
${BASE_URL}    http://localhost:8080

*** Test Cases ***
A10 SSRF - Internal Network
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    url=http://169.254.169.254/latest/meta-data/
    ${response}=    POST On Session    api    /api/v1/proxy    json=${body}    headers=${headers}    expected_status=any
    Should Be True    ${response.status_code} == 400 or ${response.status_code} == 403 or ${response.status_code} == 404

A10 SSRF - Localhost
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    url=http://localhost:5432
    ${response}=    POST On Session    api    /api/v1/proxy    json=${body}    headers=${headers}    expected_status=any
    Should Be True    ${response.status_code} == 400 or ${response.status_code} == 403 or ${response.status_code} == 404

A10 SSRF - Private IP Range
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    url=http://192.168.1.1/admin
    ${response}=    POST On Session    api    /api/v1/proxy    json=${body}    headers=${headers}    expected_status=any
    Should Be True    ${response.status_code} == 400 or ${response.status_code} == 403 or ${response.status_code} == 404

A10 SSRF - File Protocol
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    url=file:///etc/passwd
    ${response}=    POST On Session    api    /api/v1/proxy    json=${body}    headers=${headers}    expected_status=any
    Should Be True    ${response.status_code} == 400 or ${response.status_code} == 403 or ${response.status_code} == 404

A10 URL Validation Bypass
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    url=http://127.0.0.1:8080/health
    ${response}=    POST On Session    api    /api/v1/proxy    json=${body}    headers=${headers}    expected_status=any
    Should Be True    ${response.status_code} == 400 or ${response.status_code} == 403 or ${response.status_code} == 404
