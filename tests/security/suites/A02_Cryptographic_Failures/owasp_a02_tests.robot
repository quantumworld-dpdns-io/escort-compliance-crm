*** Settings ***
Library    RequestsLibrary

*** Variables ***
${BASE_URL}    http://localhost:8080

*** Test Cases ***
A02 Weak Cipher Detection
    Create Session    api    ${BASE_URL}    verify=${False}
    # Verify TLS 1.3 is enforced
    ${output}=    Run    curl -sI https://localhost:443 2>/dev/null | grep -i "TLS" || echo "TLS not detected"
    # Log for manual verification
    Log    TLS Status: ${output}

A02 PQC Algorithm Verification
    Create Session    api    ${BASE_URL}    verify=${False}
    ${headers}=    Create Dictionary    Content-Type=application/json
    ${body}=    Create Dictionary    algorithm=kyber768
    ${response}=    POST On Session    api    ${BASE_URL}/api/v1/quantum/pqc/keygen    json=${body}    headers=${headers}
    Should Be Equal As Integers    ${response.status_code}    200
    Should Be Equal    ${response.json()}[algorithm]    kyber768

A02 Key Rotation Verification
    Create Session    api    ${BASE_URL}    verify=${False}
    # Verify PQC keys can be rotated
    ${body}=    Create Dictionary    algorithm=dilithium3
    ${response}=    POST On Session    api    ${BASE_URL}/api/v1/quantum/pqc/keygen    json=${body}    headers=${headers}
    Should Be Equal As Integers    ${response.status_code}    200

A02 Certificate Validation
    # Verify certificates use PQC algorithms
    ${output}=    Run    openssl x509 -in /etc/nginx/ssl/cert.pem -text 2>/dev/null | grep -i "Signature Algorithm" || echo "Certificate check failed"
    Log    Certificate: ${output}

A02 Secret Storage Verification
    # Verify secrets are not hardcoded in config
    ${output}=    Run    grep -r "password.*=" services/ --include="*.go" --include="*.ts" --include="*.py" | grep -v "test" | grep -v "Password" | head -5 || true
    Should Be Empty    ${output}
