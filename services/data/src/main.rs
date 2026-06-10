use axum::{
    routing::get,
    Json, Router,
};
use serde::Serialize;
use tower_http::cors::{CorsLayer, Any};

#[derive(Serialize)]
struct HealthResponse {
    status: String,
}

#[derive(Serialize)]
struct DataLakeStats {
    total_files: i64,
    total_size_bytes: i64,
    file_types: Vec<String>,
}

async fn health() -> Json<HealthResponse> {
    Json(HealthResponse { status: "ok".to_string() })
}

async fn stats() -> Json<DataLakeStats> {
    Json(DataLakeStats {
        total_files: 0,
        total_size_bytes: 0,
        file_types: vec![
            "parquet".to_string(),
            "json".to_string(),
            "csv".to_string(),
        ],
    })
}

async fn query(Json(req): Json<serde_json::Value>) -> Json<serde_json::Value> {
    Json(serde_json::json!({
        "query": req.get("query").unwrap_or(&serde_json::Value::Null),
        "results": [],
        "total": 0
    }))
}

#[tokio::main]
async fn main() {
    let cors = CorsLayer::new()
        .allow_origin(Any)
        .allow_methods(Any)
        .allow_headers(Any);

    let app = Router::new()
        .route("/healthz", get(health))
        .route("/readyz", get(health))
        .route("/api/v1/data/stats", get(stats))
        .route("/api/v1/data/query", axum::routing::post(query))
        .layer(cors);

    let listener = tokio::net::TcpListener::bind("0.0.0.0:8092").await.unwrap();
    println!("Data service starting on port 8092");
    axum::serve(listener, app).await.unwrap();
}
