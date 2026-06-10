*** Settings ***
Library    RequestsLibrary

*** Variables ***
${BASE_URL}    http://localhost:8080

*** Test Cases ***
A05 Default Configuration Check
    # Verify no default passwords in production
    ${output}=    Run    grep -r "password.*admin" services/ --include="*.go" --include="*.ts" --include="*.py" --include="*.yaml" | grep -v test | grep -v example || true
    Should Be Empty    ${output}

A05 Unnecessary Services
    # Verify only required services are running
    ${output}=    Run    docker ps --format '{{.Names}}' 2>/dev/null | wc -l || echo "0"
    # Should not have excessive containers
    Should Be True    ${output} < 20

A05 Error Information Leakage
    Create Session    api    ${BASE_URL}    verify=${False}
    ${response}=    GET On Session    api    /api/v1/nonexistent    expected_status=any
    # Should not expose stack traces
    Should Not Contain    ${response.text}    stack trace
    Should Not Contain    ${response.text}    goroutine
    Should Not Contain    ${response.text}    Traceback

A05 Security Headers Verification
    Create Session    api    ${BASE_URL}    verify=${False}
    ${response}=    GET On Session    api    /health
    Dictionary Should Contain Key    ${response.headers}    X-Content-Type-Options
    Dictionary Should Contain Key    ${response.headers}    X-Frame-Options
    Dictionary Should Contain Key    ${response.headers}    X-XSS-Protection

A05 Directory Traversal
    Create Session    api    ${BASE_URL}    verify=${False}
    ${response}=    GET On Session    api    /../../../etc/passwd    expected_status=any
    Should Be True    ${response.status_code} == 400 or ${response.status_code} == 404

A05 HTTP Method Tampering
    Create Session    api    ${BASE_URL}    verify=${False}
    # Try DELETE on read-only endpoint
    ${response}=    DELETE On Session    api    /health    expected_status=any
    Should Be True    ${response.status_code} == 405 or ${response.status_code} == 404
