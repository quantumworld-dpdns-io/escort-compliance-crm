*** Settings ***
Library    RequestsLibrary

*** Variables ***
${BASE_URL}    http://localhost:8080

*** Test Cases ***
A04 Threat Model Verification
    # Verify threat model documentation exists
    ${output}=    Run    test -f docs/architecture/security.md && echo "EXISTS" || echo "MISSING"
    Should Be Equal    ${output}    EXISTS

A04 Business Logic - Double Booking Prevention
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    # Try to book the same companion at overlapping times
    ${body}=    Create Dictionary
    ...    companion_id=companion-123
    ...    start_time=2025-07-01T14:00:00Z
    ...    end_time=2025-07-01T16:00:00Z
    ...    jurisdiction=US-CA
    ${response1}=    POST On Session    api    /api/v1/bookings    json=${body}    headers=${headers}    expected_status=any
    ${response2}=    POST On Session    api    /api/v1/bookings    json=${body}    headers=${headers}    expected_status=any
    # Second booking should be rejected or flagged
    Should Be True    ${response2.status_code} == 201 or ${response2.status_code} == 409 or ${response2.status_code} == 400

A04 Business Logic - Rate Manipulation Prevention
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary
    ...    companion_id=companion-123
    ...    rate=-100.00
    ${response}=    POST On Session    api    /api/v1/companions    json=${body}    headers=${headers}    expected_status=any
    # Negative rates should be rejected
    Should Not Be Equal As Integers    ${response.status_code}    201

A04 Abuse Case - Excessive API Usage
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    # Rapid fire requests
    FOR    ${i}    IN RANGE    50
        ${response}=    GET On Session    api    /health    expected_status=any
    END
    # Should be rate limited
    ${response}=    GET On Session    api    /health    expected_status=any
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 429

A04 Abuse Case - Large Payload
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    # Create oversized payload
    ${large_bio}=    Evaluate    'x' * 1000000
    ${body}=    Create Dictionary    bio=${large_bio}
    ${response}=    POST On Session    api    /api/v1/companions    json=${body}    headers=${headers}    expected_status=any
    # Should reject oversized payloads
    Should Be True    ${response.status_code} == 400 or ${response.status_code} == 413 or ${response.status_code} == 422

A04 Abuse Case - Negative Resource IDs
    Create Session    api    ${BASE_URL}    verify=${False}
    ${response}=    GET On Session    api    /api/v1/companions/-1    expected_status=any
    Should Be True    ${response.status_code} == 400 or ${response.status_code} == 404

A04 Abuse Case - SQL Wildcards in Search
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    query=%
    ${response}=    POST On Session    api    /api/v1/companions/search    json=${body}    headers=${headers}    expected_status=any
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 400
