use chrono::Utc;
use gethostname::gethostname;
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use std::convert::Infallible;
use warp::Filter;

#[derive(Serialize, Deserialize)]
struct CalculationResponse {
    result: i32,
}

#[derive(Serialize, Deserialize)]
struct ErrorResponse {
    error: String,
}

async fn add(a: i32, b: i32) -> Result<Box<dyn warp::Reply>, Infallible> {
    let result = a + b;
    Ok(Box::new(warp::reply::json(&CalculationResponse { result })))
}

async fn add_text(a: i32, b: i32) -> Result<Box<dyn warp::Reply>, Infallible> {
    let result = a + b;
    let response = format!("a = {}, b = {}\na + b = {}\n", a, b, result);
    Ok(Box::new(warp::reply::html(response)))
}

async fn sub(a: i32, b: i32) -> Result<Box<dyn warp::Reply>, Infallible> {
    let result = a - b;
    Ok(Box::new(warp::reply::json(&CalculationResponse { result })))
}

async fn sub_text(a: i32, b: i32) -> Result<Box<dyn warp::Reply>, Infallible> {
    let result = a - b;
    let response = format!("a = {}, b = {}\na - b = {}\n", a, b, result);
    Ok(Box::new(warp::reply::html(response)))
}

async fn mul(a: i32, b: i32) -> Result<Box<dyn warp::Reply>, Infallible> {
    let result = a * b;
    Ok(Box::new(warp::reply::json(&CalculationResponse { result })))
}

async fn mul_text(a: i32, b: i32) -> Result<Box<dyn warp::Reply>, Infallible> {
    let result = a * b;
    let response = format!("a = {}, b = {}\na * b = {}\n", a, b, result);
    Ok(Box::new(warp::reply::html(response)))
}

async fn div(a: i32, b: i32) -> Result<Box<dyn warp::Reply>, Infallible> {
    if b == 0 {
        let error = ErrorResponse {
            error: "Can't divide by Zero!".to_string(),
        };
        return Ok(Box::new(warp::reply::with_status(
            warp::reply::json(&error),
            warp::http::StatusCode::BAD_REQUEST,
        )));
    }
    let result = a / b;
    Ok(Box::new(warp::reply::json(&CalculationResponse { result })))
}

async fn div_text(a: i32, b: i32) -> Result<Box<dyn warp::Reply>, Infallible> {
    if b == 0 {
        let error_response = warp::reply::with_status(
            "Can't divide by ZERO !",
            warp::http::StatusCode::BAD_REQUEST,
        );
        return Ok(Box::new(error_response));
    }

    let result = a / b;
    let response = format!("a = {}, b = {}\na / b = {}\n", a, b, result);
    Ok(Box::new(warp::reply::with_header(
        warp::reply::html(response),
        "Content-Type",
        "text/plain",
    )))
}

async fn root_handler() -> Result<Box<dyn warp::Reply>, Infallible> {
    let current_time = Utc::now().to_rfc3339();
    let hostname = gethostname()
        .into_string()
        .unwrap_or_else(|_| "unknown".to_string());

    let response = format!(
        "Hello World from Rust & Warp!\n DateTime (UTC): \"{}\"\n My hostname is \"{}\"\n",
        current_time, hostname
    );

    Ok(Box::new(warp::reply::with_header(
        warp::reply::html(response),
        "Content-Type",
        "text/plain",
    )))
}

#[tokio::main]
async fn main() {
    let port = 8080; // Порт сервера

    // Выводим сообщение о запуске сервера
    println!("Server is running on port {}...", port);

    // Root-URL ("/")
    let root_route = warp::path::end()
        .and_then(root_handler);

    // Plus
    let add_text_route = warp::path!("plus")
        .and(warp::query::<HashMap<String, i32>>())
        .and_then(|params: HashMap<String, i32>| async move {
            let a = params.get("a").unwrap_or(&0);
            let b = params.get("b").unwrap_or(&0);
            add_text(*a, *b).await
        });

    let add_route = warp::path!("api" / "plus")
        .and(warp::query::<HashMap<String, i32>>())
        .and_then(|params: HashMap<String, i32>| async move {
            let a = params.get("a").unwrap_or(&0);
            let b = params.get("b").unwrap_or(&0);
            add(*a, *b).await
        });

    // Minus
    let sub_route = warp::path!("api" / "minus")
        .and(warp::query::<HashMap<String, i32>>())
        .and_then(|params: HashMap<String, i32>| async move {
            let a = params.get("a").unwrap_or(&0);
            let b = params.get("b").unwrap_or(&0);
            sub(*a, *b).await
        });

    let sub_text_route = warp::path!("minus")
        .and(warp::query::<HashMap<String, i32>>())
        .and_then(|params: HashMap<String, i32>| async move {
            let a = params.get("a").unwrap_or(&0);
            let b = params.get("b").unwrap_or(&0);
            sub_text(*a, *b).await
        });

    // Multiply
    let mul_route = warp::path!("api" / "multiply")
        .and(warp::query::<HashMap<String, i32>>())
        .and_then(|params: HashMap<String, i32>| async move {
            let a = params.get("a").unwrap_or(&0);
            let b = params.get("b").unwrap_or(&0);
            mul(*a, *b).await
        });

    let mul_text_route = warp::path!("multiply")
        .and(warp::query::<HashMap<String, i32>>())
        .and_then(|params: HashMap<String, i32>| async move {
            let a = params.get("a").unwrap_or(&0);
            let b = params.get("b").unwrap_or(&0);
            mul_text(*a, *b).await
        });

    // Divide
    let div_route = warp::path!("api" / "divide")
        .and(warp::query::<HashMap<String, i32>>())
        .and_then(|params: HashMap<String, i32>| async move {
            let a = params.get("a").unwrap_or(&0);
            let b = params.get("b").unwrap_or(&0);
            div(*a, *b).await
        });

    let div_text_route = warp::path!("divide")
        .and(warp::query::<HashMap<String, i32>>())
        .and_then(|params: HashMap<String, i32>| async move {
            let a = params.get("a").unwrap_or(&0);
            let b = params.get("b").unwrap_or(&0);
            div_text(*a, *b).await
        });

    let routes = root_route
        .or(add_text_route)
        .or(add_route)
        .or(sub_route)
        .or(sub_text_route)
        .or(mul_route)
        .or(mul_text_route)
        .or(div_route)
        .or(div_text_route);

    warp::serve(routes).run(([0, 0, 0, 0], port)).await;
}
