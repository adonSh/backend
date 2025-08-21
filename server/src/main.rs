use axum::{Router, routing::get};
use tokio::net::TcpListener;

#[tokio::main]
async fn main() {
    let addr = String::from("0.0.0.0:8000");

    let app = Router::new()
        .route("/", get(handler))
        .route("/my_bus", get(secret_api_destination));

    let listener = TcpListener::bind(&addr).await.unwrap();

    println!("Listening on {}", &addr);

    axum::serve(listener, app).await.unwrap();
}

async fn handler() -> &'static str {
    "okie dokie artichokie"
}

async fn secret_api_destination() -> &'static str {
    "wahoo!!! u have reached my secret area! 🚌🚌🚌"
}
