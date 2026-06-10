*** Settings ***
Library    RequestsLibrary
Resource   ../../resources/api_keywords.robot

*** Test Cases ***
PQC Key Generation
    Create API Session
    ${body}=    Create Dictionary    algorithm=kyber768
    ${response}=    POST On Session    api    ${QUANTUM_URL}/pqc/keygen    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200
    Dictionary Should Contain Key    ${response.json()}    public_key
    Dictionary Should Contain Key    ${response.json()}    private_key
    Should Be Equal    ${response.json()}[algorithm]    kyber768

PQC Encryption
    Create API Session
    ${body}=    Create Dictionary
    ...    public_key=test-pub-key
    ...    plaintext=Sensitive compliance data
    ${response}=    POST On Session    api    ${QUANTUM_URL}/pqc/encrypt    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200
    Dictionary Should Contain Key    ${response.json()}    ciphertext

PQC Decryption
    Create API Session
    ${body}=    Create Dictionary
    ...    private_key=test-priv-key
    ...    ciphertext=test-ciphertext
    ${response}=    POST On Session    api    ${QUANTUM_URL}/pqc/decrypt    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200
    Dictionary Should Contain Key    ${response.json()}    plaintext

PQC Digital Signature
    Create API Session
    ${body}=    Create Dictionary
    ...    private_key=test-priv-key
    ...    message=Important compliance document
    ${response}=    POST On Session    api    ${QUANTUM_URL}/pqc/sign    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200
    Dictionary Should Contain Key    ${response.json()}    signature

PQC Signature Verification
    Create API Session
    ${body}=    Create Dictionary
    ...    public_key=test-pub-key
    ...    message=Important compliance document
    ...    signature=test-signature
    ${response}=    POST On Session    api    ${QUANTUM_URL}/pqc/verify    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200
    Should Be Equal    ${response.json()}[valid]    ${TRUE}

Quantum Optimization
    Create API Session
    ${body}=    Create Dictionary
    ...    algorithm=qaoa
    ...    cost_function=booking_scheduling
    ...    variables=10
    ${response}=    POST On Session    api    ${QUANTUM_URL}/optimize    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200
    Dictionary Should Contain Key    ${response.json()}    solution
    Should Be Equal    ${response.json()}[quantum_used]    ${TRUE}

Quantum ML Prediction
    Create API Session
    ${body}=    Create Dictionary
    ...    model=safety_scorer
    ...    features=${{"age": 25, "jurisdiction": "US-CA", "history": "clean"}}
    ${response}=    POST On Session    api    ${QUANTUM_URL}/ml/predict    json=${body}    headers=${HEADERS}
    Should Be Equal As Integers    ${response.status_code}    200
    Dictionary Should Contain Key    ${response.json()}    prediction
    Dictionary Should Contain Key    ${response.json()}    confidence

QRNG Random Bytes
    Create API Session
    ${response}=    GET On Session    api    ${QUANTUM_URL}/qrng/random
    Should Be Equal As Integers    ${response.status_code}    200
    Dictionary Should Contain Key    ${response.json()}    random_bytes
    Should Be Equal    ${response.json()}[source]    qrng
