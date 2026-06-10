*** Settings ***
Library    RequestsLibrary
Resource   ../../resources/api_keywords.robot

*** Test Cases ***
Query Data Lake
    Create API Session
    ${body}=    Create Dictionary
    ...    query=SELECT count(*) FROM bookings WHERE jurisdiction = 'US-CA'
    ${response}=    POST On Session    api    ${BASE_URL}/api/v1/data/query    json=${body}    headers=${HEADERS}
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 201

List Iceberg Tables
    Create API Session
    ${response}=    GET On Session    api    ${BASE_URL}/api/v1/data/tables    headers=${HEADERS}
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 201

Data Ingestion
    Create API Session
    ${body}=    Create Dictionary
    ...    source=bookings
    ...    format=parquet
    ${response}=    POST On Session    api    ${BASE_URL}/api/v1/data/ingest    json=${body}    headers=${HEADERS}
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 201

Data Lineage
    Create API Session
    ${response}=    GET On Session    api    ${BASE_URL}/api/v1/data/lineage    headers=${HEADERS}
    Should Be True    ${response.status_code} == 200 or ${response.status_code} == 201
