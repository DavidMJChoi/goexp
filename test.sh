count=1
while [[ $count -le 10 ]]; do
    go run src/goplch9/pingpong/pingpong.go
    ((count++))
done