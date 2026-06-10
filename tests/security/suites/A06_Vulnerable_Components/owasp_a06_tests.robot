*** Settings ***
Library    RequestsLibrary
Library    OperatingSystem

*** Variables ***
${BASE_URL}    http://localhost:8080

*** Test Cases ***
A06 Known Vulnerability Scan - Go
    ${output}=    Run    cd services/shared && go list -m -json all 2>/dev/null | grep -i "vulnerability\|CVE" || true
    Log    Go dependencies: ${output}

A06 Known Vulnerability Scan - Node
    ${output}=    Run    cd services/payments && npm audit 2>/dev/null | grep -i "high\|critical" || true
    Log    Node dependencies: ${output}

A06 Known Vulnerability Scan - Python
    ${output}=    Run    cd services/ml && pip-audit 2>/dev/null | grep -i "high\|critical" || true
    Log    Python dependencies: ${output}

A06 Known Vulnerability Scan - Rust
    ${output}=    Run    cd services/crypto && cargo audit 2>/dev/null | grep -i "high\|critical" || true
    Log    Rust dependencies: ${output}

A06 Outdated Component Detection
    ${output}=    Run    go list -u -m all 2>/dev/null | grep -i "\->" | head -10 || true
    Log    Outdated Go modules: ${output}

A06 Container Image Scan
    ${output}=    Run    docker images escort-crm/* --format "{{.Repository}}:{{.Tag}}" 2>/dev/null | head -5 || true
    Log    Container images: ${output}
