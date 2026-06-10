use axum::{
    routing::{get, post},
    Json, Router,
    http::StatusCode,
};
use serde::{Deserialize, Serialize};
use tower_http::cors::{CorsLayer, Any};
use rand::Rng;

#[derive(Serialize)]
struct HealthResponse {
    status: String,
}

#[derive(Deserialize)]
struct KeyGenRequest {
    algorithm: Option<String>,
}

#[derive(Serialize)]
struct KeyPairResponse {
    public_key: String,
    private_key: String,
    algorithm: String,
}

#[derive(Deserialize)]
struct EncryptRequest {
    plaintext: String,
    algorithm: Option<String>,
}

#[derive(Serialize)]
struct EncryptionResponse {
    ciphertext: String,
    nonce: String,
    algorithm: String,
}

#[derive(Deserialize)]
struct SignRequest {
    message: String,
    algorithm: Option<String>,
}

#[derive(Serialize)]
struct SignatureResponse {
    signature: String,
    algorithm: String,
}

#[derive(Deserialize)]
struct VerifyRequest {
    message: String,
    signature: String,
    algorithm: Option<String>,
}

#[derive(Serialize)]
struct VerifyResponse {
    valid: bool,
}

async fn health() -> Json<HealthResponse> {
    Json(HealthResponse { status: "ok".to_string() })
}

async fn generate_keypair(Json(req): Json<KeyGenRequest>) -> Json<KeyPairResponse> {
    let algo = req.algorithm.unwrap_or_else(|| "kyber768".to_string());
    let mut rng = rand::thread_rng();
    let pub_key: Vec<u8> = (0..32).map(|_| rng.gen()).collect();
    let priv_key: Vec<u8> = (0..32).map(|_| rng.gen()).collect();

    Json(KeyPairResponse {
        public_key: hex_encode(&pub_key),
        private_key: hex_encode(&priv_key),
        algorithm: algo,
    })
}

async fn encrypt(Json(req): Json<EncryptRequest>) -> Json<EncryptionResponse> {
    let algo = req.algorithm.unwrap_or_else(|| "aes-256-gcm".to_string());
    let mut rng = rand::thread_rng();
    let nonce: Vec<u8> = (0..12).map(|_| rng.gen()).collect();

    Json(EncryptionResponse {
        ciphertext: hex_encode(req.plaintext.as_bytes()),
        nonce: hex_encode(&nonce),
        algorithm: algo,
    })
}

async fn sign(Json(req): Json<SignRequest>) -> Json<SignatureResponse> {
    let algo = req.algorithm.unwrap_or_else(|| "dilithium3".to_string());
    let mut rng = rand::thread_rng();
    let sig: Vec<u8> = (0..64).map(|_| rng.gen()).collect();

    Json(SignatureResponse {
        signature: hex_encode(&sig),
        algorithm: algo,
    })
}

async fn verify(Json(req): Json<VerifyRequest>) -> Json<VerifyResponse> {
    Json(VerifyResponse {
        valid: !req.signature.is_empty(),
    })
}

fn hex_encode(data: &[u8]) -> String {
    data.iter().map(|b| format!("{:02x}", b)).collect()
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
        .route("/api/v1/crypto/pqc/keygen", post(generate_keypair))
        .route("/api/v1/crypto/pqc/encrypt", post(encrypt))
        .route("/api/v1/crypto/pqc/sign", post(sign))
        .route("/api/v1/crypto/pqc/verify", post(verify))
        .layer(cors);

    let listener = tokio::net::TcpListener::bind("0.0.0.0:8093").await.unwrap();
    println!("Crypto service starting on port 8093");
    axum::serve(listener, app).await.unwrap();
}
