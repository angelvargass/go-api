NUMBER_OF_REQUESTS=100
URL="localhost:8080/v1/sample"

make_api_call() {
    call_number=$1
    echo "Making API call $call_number"
    response=$(curl -s "$URL")
    echo "Response: $response"
}

make_multiple_api_calls() {
    for ((i = 0; i < NUMBER_OF_REQUESTS; i++)); do
        make_api_call $((i+1)) &
    done
    wait
}

main() {
    start_time=$(gdate +%s%3N)
    make_multiple_api_calls
    end_time=$(gdate +%s%3N)
    elapsed_time=$((end_time - start_time))
    echo "Script completed in ${elapsed_time} ms"
}

main