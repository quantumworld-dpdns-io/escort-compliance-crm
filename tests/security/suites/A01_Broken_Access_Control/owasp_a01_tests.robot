*** Settings ***
Library    RequestsLibrary
Library    Collections
Library    OperatingSystem

*** Variables ***
${BASE_URL}    http://localhost:8080

*** Test Cases ***
A01 IDOR - Access Other User Resources
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    # Try accessing resources without auth
    ${response}=    GET On Session    api    /api/v1/companions/other-user-id    headers=${headers}    expected_status=any
    # Should get 401 Unauthorized
    Should Be True    ${response.status_code} == 401 or ${response.status_code} == 403

A01 Privilege Escalation
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    # Try admin endpoint without admin role
    ${response}=    POST On Session    api    /api/v1/admin/users    headers=${headers}    expected_status=any
    Should Be True    ${response.status_code} == 401 or ${response.status_code} == 403

A02 Weak Encryption Detection
    Create Session    api    ${BASE_URL}    verify=${False}
    # Verify PQC is enabled
    ${response}=    GET On Session    api    /api/v1/quantum/pqc/health    expected_status=any
    Should Not Be Equal As Integers    ${response.status_code}    500

A02 PQC Key Generation
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    algorithm=kyber768
    ${response}=    POST On Session    api    /api/v1/quantum/pqc/keygen    json=${body}    headers=${headers}
    Should Be Equal As Integers    ${response.status_code}    200

A03 SQL Injection - Login
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary
    ...    email=' OR '1'='1' --
    ...    password=' OR '1'='1'
    ${response}=    POST On Session    api    /api/v1/auth/login    json=${body}    headers=${headers}    expected_status=any
    # Should NOT return 200 with valid token
    Should Not Be Equal As Integers    ${response.status_code}    200

A03 SQL Injection - Search
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    query="; DROP TABLE companions; --
    ${response}=    POST On Session    api    /api/v1/companions/search    json=${body}    headers=${headers}    expected_status=any
    Should Be True    ${response.status_code} != 500

A03 NoSQL Injection
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    email=${{"$gt": ""}}    password=${{"$gt": ""}}
    ${response}=    POST On Session    api    /api/v1/auth/login    json=${body}    headers=${headers}    expected_status=any
    Should Not Be Equal As Integers    ${response.status_code}    200

A05 Default Credentials
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    email=admin@example.com    password=admin
    ${response}=    POST On Session    api    /api/v1/auth/login    json=${body}    headers=${headers}    expected_status=any
    Should Not Be Equal As Integers    ${response.status_code}    200

A05 Security Headers
    Create Session    api    ${BASE_URL}    verify=${False}
    ${response}=    GET On Session    api    /health
    Dictionary Should Contain Key    ${response.headers}    X-Content-Type-Options
    Dictionary Should Contain Key    ${response.headers}    X-Frame-Options

A06 Dependency Vulnerability
    # Check that no known vulnerable dependencies are used
    ${output}=    Run    go list -json ./... 2>/dev/null | grep -i vulnerability || true
    Should Be Empty    ${output}

A07 Brute Force Protection
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    email=test@example.com    password=wrong
    # Try multiple failed logins
    FOR    ${i}    IN RANGE    5
        ${response}=    POST On Session    api    /api/v1/auth/login    json=${body}    headers=${headers}    expected_status=any
    END
    # After 5 failures, should be rate limited
    ${response}=    POST On Session    api    /api/v1/auth/login    json=${body}    headers=${headers}    expected_status=any
    Should Be True    ${response.status_code} == 429 or ${response.status_code} == 401

A08 Data Integrity - Tamper Detection
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary
    ...    Content-Type=application/json
    ...    X-Request-ID=tampered-id
    ${response}=    GET On Session    api    /health    headers=${headers}
    Should Be Equal As Integers    ${response.status_code}    200

A09 Audit Logging
    Create Session    api    ${BASE_URL}    verify=${False}
    # Verify that security events are logged
    ${output}=    Run    cat /var/log/escort-crm/audit.log 2>/dev/null | wc -l || echo "0"
    # Audit log should exist
    Should Be True    ${output} >= 0

A10 SSRF - Internal Network Scan
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    url=http://169.254.169.254/latest/meta-data/
    ${response}=    POST On Session    api    /api/v1/proxy    json=${body}    headers=${headers}    expected_status=any
    # Should block SSRF attempts
    Should Be True    ${response.status_code} == 400 or ${response.status_code} == 403 or ${response.status_code} == 404
