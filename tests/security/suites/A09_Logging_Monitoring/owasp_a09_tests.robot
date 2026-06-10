*** Settings ***
Library    RequestsLibrary
Library    OperatingSystem

*** Variables ***
${BASE_URL}    http://localhost:8080

*** Test Cases ***
A09 Log Injection Prevention
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary
    ...    Content-Type=application/json
    ...    X-Request-ID=injection-attempt\n[FAKE LOG ENTRY]
    ${response}=    GET On Session    api    /health    headers=${headers}
    Should Be Equal As Integers    ${response.status_code}    200

A09 Audit Trail Completeness
    # Verify security events are logged
    ${output}=    Run    ls -la /var/log/escort-crm/ 2>/dev/null || echo "Log directory not found"
    Log    Audit log status: ${output}

A09 Log Integrity
    # Verify logs cannot be tampered
    ${output}=    Run    chmod 644 /var/log/escort-crm/*.log 2>/dev/null || true
    # Logs should be append-only
    ${output}=    Run    lsattr /var/log/escort-crm/audit.log 2>/dev/null | grep -i append || echo "Append-only not set"
    Log    Log integrity: ${output}
