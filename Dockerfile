FROM rust:1.89

WORKDIR /app

COPY server/  ./

RUN cargo build --release

RUN mv target/release/server .

RUN find . -maxdepth 1 -type f ! -name "server" -exec rm {} +

EXPOSE 8000

CMD ["./server"]
