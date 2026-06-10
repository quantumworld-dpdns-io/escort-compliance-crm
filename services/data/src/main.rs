use axum::{routing::get, Json, Router};
use serde::{Deserialize, Serialize};
use tower_http::cors::CorsLayer;
use tracing::info;

#[derive(Serialize, Deserialize)]
struct HealthResponse {
    status: String,
    service: String,
    data_lake_enabled: bool,
}

async fn health() -> Json<HealthResponse> {
    Json(HealthResponse {
        status: "healthy".to_string(),
        service: "data".to_string(),
        data_lake_enabled: true,
    })
}

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt::init();
    info!("Starting Data Lakehouse Service");

    let app = Router::new()
        .route("/health", get(health))
        .layer(CorsLayer::permissive());

    let listener = tokio::net::TcpListener::bind("0.0.0.0:8092")
        .await
        .unwrap();
    info!("Data service listening on :8092");
    axum::serve(listener, app).await.unwrap();
}
